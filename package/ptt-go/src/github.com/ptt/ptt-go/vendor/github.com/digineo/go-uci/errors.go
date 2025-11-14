package uci

import "fmt"

// ErrConfigAlreadyLoaded is returned by LoadConfig, if the given config
// name is already present.
type ErrConfigAlreadyLoaded struct {
	Name string
}

func (err ErrConfigAlreadyLoaded) Error() string {
	return fmt.Sprintf("%s already loaded", err.Name)
}

// ErrUnknownOptionType is returned when trying to parse an invalid OptionType.
type ErrUnknownOptionType struct {
	Type string
}

func (err ErrUnknownOptionType) Error() string {
	return fmt.Sprintf("Unknown Option type %s", err.Type)
}

// IsConfigAlreadyLoaded reports, whether err is of type ErrConfigAlredyLoaded.
//
// Deprecated: use errors.Is or errors.As.
func IsConfigAlreadyLoaded(err error) bool {
	if err == nil {
		return false
	}
	_, is := err.(*ErrConfigAlreadyLoaded) //nolint:errorlint
	return is
}

// ErrSectionTypeMismatch is returned by AddSection if the section-to-add
// already exists with a different type.
type ErrSectionTypeMismatch struct {
	Config, Section string // name
	ExistingType    string
	NewType         string
}

func (err ErrSectionTypeMismatch) Error() string {
	return fmt.Sprintf("type mismatch for %s.%s, got %s, want %s",
		err.Config, err.Section, err.ExistingType, err.NewType)
}

// IsSectionTypeMismatch reports, whether err is of type ErrSectionTypeMismatch.
//
// Deprecated: use errors.Is or errors.As.
func IsSectionTypeMismatch(err error) bool {
	if err == nil {
		return false
	}
	_, is := err.(*ErrSectionTypeMismatch) //nolint:errorlint
	return is
}

type ParseError string

func (err ParseError) Error() string {
	return fmt.Sprintf("parse error: %s", string(err))
}

// IsParseError reports, whether err is of type ParseError.
//
// Deprecated: use errors.Is or errors.As.
func IsParseError(err error) bool {
	if err == nil {
		return false
	}
	_, is := err.(*ParseError) //nolint:errorlint
	return is
}
