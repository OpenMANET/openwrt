package uci

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strconv"
	"strings"
)

// NOTE: config, section and option types basically are AST nodes for the
// parser. The JSON struct tags are mainly for development and testing
// purposes: We'er generating JSON dumps of the tree when running tests
// with DUMP="json". After a manual comparison with the corresponding UCI
// file in testdata/, we can use the dumps to read them back as test case
// expectations.

// config represents a file in UCI. It consists of sections.
type config struct {
	Name     string     `json:"name"`
	Sections []*section `json:"sections,omitempty"`

	tainted bool // changed by tree methods when things were modified
}

// newConfig returns a new config object.
func newConfig(name string) *config {
	return &config{
		Name:     name,
		Sections: make([]*section, 0, 1),
	}
}

func (c *config) WriteTo(w io.Writer) (n int64, err error) {
	var buf bytes.Buffer

	for _, sec := range c.Sections {
		if sec.Name == "" {
			fmt.Fprintf(&buf, "\nconfig %s\n", sec.Type)
		} else {
			fmt.Fprintf(&buf, "\nconfig %s '%s'\n", sec.Type, sec.Name)
		}

		for _, opt := range sec.Options {
			switch opt.Type {
			case TypeOption:
				fmt.Fprintf(&buf, "\toption %s '%s'\n", opt.Name, opt.Values[0])
			case TypeList:
				for _, v := range opt.Values {
					fmt.Fprintf(&buf, "\tlist %s '%s'\n", opt.Name, v)
				}
			}
		}
	}
	buf.WriteByte('\n')
	return buf.WriteTo(w)
}

// Get fetches a section by name.
//
// Support for unnamed section notation (@foo[idx]) is present.
func (c *config) Get(name string) *section {
	if strings.HasPrefix(name, "@") {
		sec, _ := c.getUnnamed(name) // TODO: log error?
		return sec
	}
	return c.getNamed(name)
}

func (c *config) getNamed(name string) *section {
	for _, sec := range c.Sections {
		if sec.Name == name {
			return sec
		}
	}
	return nil
}

var (
	ErrImplausibleSectionSelector = errors.New("implausible section selector: must be at least 5 characters long")
	ErrMustStartWithAt            = errors.New("invalid syntax: section selector must start with @ sign")
	ErrMultipleAtSigns            = errors.New("invalid syntax: multiple @ signs found")
	ErrMultipleOpenBrackets       = errors.New("invalid syntax: multiple open brackets found")
	ErrMultipleCloseBrackets      = errors.New("invalid syntax: multiple closed brackets found")
	ErrInvalidSectionSelector     = errors.New("invalid syntax: section selector must have format '@type[index]'")
)

func unmangleSectionName(name string) (typ string, index int, err error) { //nolint:cyclop
	l := len(name)
	if l < 5 { // "@a[0]"
		err = ErrImplausibleSectionSelector
		return
	}
	if name[0] != '@' {
		err = ErrMustStartWithAt
		return
	}

	bra, ket := 0, l-1 // bracket positions
	for i, r := range name {
		switch {
		case i != 0 && r == '@':
			err = ErrMultipleAtSigns
			return
		case r == '[' && bra > 0:
			err = ErrMultipleOpenBrackets
			return
		case r == ']' && i != ket:
			err = ErrMultipleCloseBrackets
			return
		case r == '[':
			bra = i
		}
	}

	if bra == 0 || bra >= ket {
		err = ErrInvalidSectionSelector
		return
	}

	typ = name[1:bra]
	index, err = strconv.Atoi(name[bra+1 : ket])
	if err != nil {
		err = fmt.Errorf("invalid syntax: index must be numeric: %w", err)
	}
	return typ, index, err
}

var ErrUnnamedIndexOutOfBounds = errors.New("invalid name: index out of bounds")

func (c *config) getUnnamed(name string) (*section, error) {
	typ, idx, err := unmangleSectionName(name)
	if err != nil {
		return nil, err
	}

	count := c.count(typ)
	if -count > idx || idx >= count {
		return nil, ErrUnnamedIndexOutOfBounds
	}
	if idx < 0 {
		idx += count // count from the end
	}

	for i, n := 0, 0; i < len(c.Sections); i++ {
		if c.Sections[i].Type == typ {
			if idx == n {
				return c.Sections[i], nil
			}
			n++
		}
	}
	return nil, nil
}

func (c *config) Add(s *section) *section {
	c.Sections = append(c.Sections, s)
	return s
}

func (c *config) Merge(s *section) *section {
	var sec *section
	for i := range c.Sections {
		sname := c.sectionName(s)
		cname := c.sectionName(c.Sections[i])

		if sname == cname {
			sec = c.Sections[i]
			break
		}
	}

	if sec == nil {
		return c.Add(s)
	}
	for _, o := range s.Options {
		sec.Merge(o)
	}
	return sec
}

func (c *config) Del(name string) {
	var i int
	for i = 0; i < len(c.Sections); i++ {
		if c.Sections[i].Name == name {
			break
		}
	}
	if i < len(c.Sections) {
		c.Sections = append(c.Sections[:i], c.Sections[i+1:]...)
	}
}

func (c *config) sectionName(s *section) string {
	if s.Name != "" {
		return s.Name
	}
	return fmt.Sprintf("@%s[%d]", s.Type, c.index(s))
}

func (c *config) index(s *section) (i int) {
	for _, sec := range c.Sections {
		if sec == s {
			return i
		}
		if sec.Type == s.Type {
			i++
		}
	}
	panic("not reached")
}

func (c *config) count(typ string) (n int) {
	for _, sec := range c.Sections {
		if sec.Type == typ {
			n++
		}
	}
	return
}

// A section represents a group of options in UCI. It may be named or
// unnamed. In the latter case, its synthetic name is constructed from
// the section type and index (e.g. "@system[0]").
type section struct {
	Name    string    `json:"name,omitempty"`
	Type    string    `json:"type"`
	Options []*option `json:"options,omitempty"`
}

// newSection returns a new section object.
func newSection(typ, name string) *section {
	return &section{
		Type:    typ,
		Name:    name,
		Options: make([]*option, 0, 1),
	}
}

func (s *section) Add(o *option) {
	s.Options = append(s.Options, o)
}

func (s *section) Merge(o *option) {
	for _, opt := range s.Options {
		if opt.Name == o.Name {
			opt.MergeValues(o.Values...)
			return
		}
	}
	s.Options = append(s.Options, o)
}

// Del removes an option with the given name. It returns whether the
// option actually existed.
func (s *section) Del(name string) bool {
	var i int
	for i = 0; i < len(s.Options); i++ {
		if s.Options[i].Name == name {
			break
		}
	}

	if i == len(s.Options) {
		return false
	}

	s.Options = append(s.Options[:i], s.Options[i+1:]...)

	return true
}

// Get fetches an option by name.
func (s *section) Get(name string) *option {
	for _, opt := range s.Options {
		if opt.Name == name {
			return opt
		}
	}
	return nil
}

// An Option is the key to one or more values. Multiple values indicate
// a list option.
type option struct {
	Name   string     `json:"name"`
	Values []string   `json:"values"`
	Type   OptionType `json:"type"`
}

// newOption returns a new option object.
func newOption(name string, optionType OptionType, values ...string) *option {
	return &option{
		Name:   name,
		Values: values,
		Type:   optionType,
	}
}

func (o *option) SetValues(vs ...string) {
	o.Values = vs
}

func (o *option) AddValue(v string) {
	o.Values = append(o.Values, v)
}

func (o *option) MergeValues(vs ...string) {
	have := make(map[string]struct{})
	for _, v := range o.Values {
		have[v] = struct{}{}
	}

	for _, v := range vs {
		if _, exists := have[v]; exists {
			continue
		}
		o.AddValue(v)
	}
}
