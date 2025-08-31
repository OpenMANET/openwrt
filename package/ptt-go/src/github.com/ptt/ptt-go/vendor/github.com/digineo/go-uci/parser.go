package uci

import (
	"fmt"
)

// scanner is intertwined with lexer and groups lexemes into token
// (which is a typed list of items).
//
// Refer to the implementation of lexer for hints about the design.
// The scanner is strongly modeled after the same principles, although
// a bit less elegant at times.
type scanner struct {
	lexer  *lexer
	state  scanFn
	last   *item  // last item read from the lexer, but deffered by the state
	curr   []item // accepted items
	tokens chan token
}

func scan(name, input string) *scanner {
	return &scanner{
		lexer:  lex(name, input),
		state:  scanStart,
		curr:   make([]item, 0, 3),
		tokens: make(chan token, 2),
	}
}

func (s *scanner) nextToken() token {
	for s.state != nil {
		select {
		case tok, ok := <-s.tokens:
			if ok {
				return tok
			}
			return s.eof()
		default:
			st := s.state(s)
			s.state = st
			if st == nil {
				return s.stop()
			}
		}
	}
	return s.eof()
}

func (s *scanner) eof() token {
	return token{typ: tokEOF}
}

func (s *scanner) stop() token {
	tok := s.eof()
	if s.tokens == nil {
		return tok
	}
	s.lexer.stop()
	if len(s.tokens) > 0 {
		tok = <-s.tokens
	}
	close(s.tokens)
	s.tokens = nil
	return tok
}

func (s *scanner) next() item {
	if s.last != nil {
		it := *s.last
		s.last = nil
		return it
	}
	return s.lexer.nextItem()
}

func (s *scanner) peek() item {
	it := s.next()
	s.backup(it)
	return it
}

func (s *scanner) backup(it item) {
	s.last = &it
}

func (s *scanner) accept(it itemType) bool {
	tok := s.next()
	if tok.typ == it {
		s.curr = append(s.curr, tok)
		return true
	}
	s.backup(tok)
	return false
}

func (s *scanner) emit(typ scanToken) {
	s.tokens <- token{typ: typ, items: s.curr}
	s.curr = make([]item, 0, 3)
}

func (s *scanner) errorf(format string, args ...interface{}) scanFn {
	s.tokens <- token{
		typ:   tokError,
		items: []item{{itemError, fmt.Sprintf(format, args...), 0}},
	}
	return nil
}

// scanStart looks for a "package" or "config" item.
func scanStart(s *scanner) scanFn {
	switch it := s.next(); it.typ { //nolint:exhaustive
	case itemPackage:
		return scanPackage
	case itemConfig:
		return scanSection
	case itemError:
		return s.errorf(it.val)
	case itemEOF:
		return nil
	default:
		return s.errorf("expected package or config token, got %s", it)
	}
}

// scanPackage looks for a package name.
func scanPackage(s *scanner) scanFn {
	switch it := s.next(); it.typ { //nolint:exhaustive
	case itemString:
		s.curr = append(s.curr, it)
		s.emit(tokPackage)
		return scanStart
	case itemError:
		return s.errorf(it.val)
	default:
		return s.errorf("expected string value while parsing package, got %s", it)
	}
}

// scanSection looks for a section type and optional a name.
func scanSection(s *scanner) scanFn {
	switch it := s.next(); it.typ { //nolint:exhaustive
	case itemIdent:
		s.curr = append(s.curr, it)
		// the name is optional
		if tok := s.peek(); tok.typ == itemString {
			s.accept(itemString)
		}
		s.emit(tokSection)
		return scanOption
	case itemError:
		return s.errorf(it.val)
	default:
		return s.errorf("expected identifier while parsing config section, got %s", it)
	}
}

// scanOption looks for either an "option" or "list" keyword (with name
// and value), or it falls back to scanStart.
func scanOption(s *scanner) scanFn {
	it := s.next()
	switch it.typ { //nolint:exhaustive
	case itemOption:
		return scanOptionName
	case itemList:
		return scanListName
	case itemError:
		return s.errorf(it.val)
	default:
		s.backup(it)
		return scanStart
	}
}

// scanOptionName looks for a name of a string option.
func scanOptionName(s *scanner) scanFn {
	if s.accept(itemIdent) {
		return scanOptionValue
	}
	return s.errorf("expected option name")
}

// scanListName looks for a name of a list option.
func scanListName(s *scanner) scanFn {
	if s.accept(itemIdent) {
		return scanListValue
	}
	return s.errorf("expected option name")
}

// scanOptionValue looks for the value associated with an option.
func scanOptionValue(s *scanner) scanFn {
	switch it := s.next(); it.typ { //nolint:exhaustive
	case itemString:
		s.curr = append(s.curr, it)
		s.emit(tokOption)
		return scanOption
	case itemError:
		return s.errorf(it.val)
	default:
		return s.errorf("expected option value, got %s", it)
	}
}

// scanListValue looks for the value associated with an option.
func scanListValue(s *scanner) scanFn {
	switch it := s.next(); it.typ { //nolint:exhaustive
	case itemString:
		s.curr = append(s.curr, it)
		s.emit(tokList)
		return scanOption
	case itemError:
		return s.errorf(it.val)
	default:
		return s.errorf("expected option value, got %s", it)
	}
}

func (s *scanner) each(fn func(token) bool) bool {
	for tok := s.nextToken(); tok.typ != tokEOF; tok = s.nextToken() {
		if !fn(tok) {
			s.stop()
			return false
		}
	}
	return true
}

// parse tries to parse a named input string into a config object.
func parse(name, input string) (cfg *config, err error) {
	cfg = newConfig(name)
	var sec *section

	scan(name, input).each(func(tok token) bool {
		switch tok.typ { //nolint:exhaustive
		case tokError:
			perr := ParseError(tok.items[0].val)
			err = &perr
			return false

		case tokPackage:
			err = ParseError("UCI imports/exports are not yet supported")
			return false

		case tokSection:
			name := tok.items[0].val
			if len(tok.items) == 2 {
				sec = cfg.Merge(newSection(name, tok.items[1].val))
			} else {
				sec = cfg.Add(newSection(name, ""))
			}

		case tokOption:
			name := tok.items[0].val
			val := tok.items[1].val

			if opt := sec.Get(name); opt != nil {
				opt.SetValues(val)
			} else {
				sec.Add(newOption(name, TypeOption, val))
			}

		case tokList:
			name := tok.items[0].val
			val := tok.items[1].val

			if opt := sec.Get(name); opt != nil {
				opt.MergeValues(val)
			} else {
				sec.Add(newOption(name, TypeList, val))
			}
		}
		return true
	})
	return cfg, err
}
