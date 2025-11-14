package uci

import (
	"bytes"
	"fmt"
)

// item represents a lexeme (token)
//
// https://talks.golang.org/2011/lex.slide#8
type item struct {
	typ itemType
	val string
	pos int
}

type OptionType int

const (
	TypeOption OptionType = iota // option is not a list
	TypeList                     // option is a list
)

// MarshalJSON implements encoding/json.Marshaler.
func (ot OptionType) MarshalJSON() ([]byte, error) {
	switch ot {
	case TypeOption:
		return []byte(`"option"`), nil
	case TypeList:
		return []byte(`"list"`), nil
	default:
		return nil, ErrUnknownOptionType{Type: fmt.Sprintf("!OptionType(%02x)", ot)}
	}
}

// UnmarshalJSON implements encoding/json.Unmarshaler.
func (ot *OptionType) UnmarshalJSON(b []byte) error {
	if len(b) == 0 || bytes.Equal(b, []byte("null")) || bytes.Equal(b, []byte("\"option\"")) {
		*ot = TypeOption
		return nil
	}
	if bytes.Equal(b, []byte(`"list"`)) {
		*ot = TypeList
		return nil
	}
	return ErrUnknownOptionType{Type: string(b)}
}

// itemType defines the kind of lexed item
//
// https://talks.golang.org/2011/lex.slide#9
type itemType int

// These items define the UCI language.
const (
	itemError itemType = iota // error occurred; item.val is text of error

	itemBOF // begin of file; lexing starts here
	itemEOF // end of file; lexing ends here

	itemPackage // package keyword
	itemConfig  // config keyword
	itemOption  // option keyword
	itemList    // list keyword
	itemIdent   // identifier string
	itemString  // quoted string
)

func (t itemType) String() string {
	switch t {
	case itemError:
		return "Error"
	case itemBOF:
		return "BOF"
	case itemEOF:
		return "EOF"
	case itemPackage:
		return "Package"
	case itemConfig:
		return "Config"
	case itemOption:
		return "Option"
	case itemList:
		return "List"
	case itemIdent:
		return "Ident"
	case itemString:
		return "String"
	}
	return fmt.Sprintf("%%itemType(%d)", int(t))
}

// keyword represents a special marker of the input: each (trimmed,
// non-empty) line of the input must start with a keywords.
type keyword string

// these are the recognized keywords.
const (
	kwPackage = keyword("package")
	kwConfig  = keyword("config")
	kwOption  = keyword("option")
	kwList    = keyword("list")
)

// String implements fmt.Stringer interface. Useful for debugging
//
// https://talks.golang.org/2011/lex.slide#11
func (i item) String() string {
	if i.pos < 0 {
		if i.typ != itemError && len(i.val) > 25 {
			return fmt.Sprintf("(%s %.25q...)", i.typ, i.val)
		}
		return fmt.Sprintf("(%s %q)", i.typ, i.val)
	}

	if i.typ != itemError && len(i.val) > 25 {
		return fmt.Sprintf("(%s %.25q... %d)", i.typ, i.val, i.pos)
	}
	return fmt.Sprintf("(%s %q %d)", i.typ, i.val, i.pos)
}

type scanFn func(*scanner) scanFn

type scanToken int

const (
	tokError scanToken = iota
	tokEOF

	tokPackage // item-seq: (package, string)
	tokSection // item-seq: (config, ident, maybe string)
	tokOption  // item-seq: (option, ident, string)
	tokList    // item-seq: (list, ident, string)
)

func (t scanToken) String() string {
	switch t {
	case tokEOF:
		return "eof"
	case tokError:
		return "error"
	case tokPackage:
		return "package"
	case tokSection:
		return "config"
	case tokOption:
		return "option"
	case tokList:
		return "list"
	}
	return fmt.Sprintf("%%scanToken(%d)", int(t))
}

type token struct {
	typ   scanToken
	items []item
}

func (t token) String() string {
	return fmt.Sprintf("%s%s", t.typ, t.items)
}
