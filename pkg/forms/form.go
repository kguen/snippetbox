package forms

import (
	"fmt"
	"net/url"
	"regexp"
	"strings"
	"unicode/utf8"
)

type Form struct {
	url.Values
	Errors errors
}

var EmailRX = regexp.MustCompile("^[a-zA-Z0-9.!#$%&'*+\\/=?^_`{|}~-]+@[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?(?:\\.[a-zA-Z0-9](?:[a-zA-Z0-9-]{0,61}[a-zA-Z0-9])?)*$")

func New(data url.Values) *Form {
	return &Form{data, errors(make(map[string][]string))}
}

func (f *Form) Required(fields ...string) {
	for _, name := range fields {
		if strings.TrimSpace(f.Get(name)) == "" {
			f.Errors.Add(name, "This field cannot be blank")
		}
	}
}

func (f *Form) MaxLength(field string, length int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) > length {
		f.Errors.Add(field, fmt.Sprintf("This field is too long (maximum is %d characters)", length))
	}
}

func (f *Form) MinLength(field string, length int) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if utf8.RuneCountInString(value) < length {
		f.Errors.Add(field, fmt.Sprintf("This field is too short (minimum is %d characters)", length))
	}
}

func (f *Form) PermittedValues(field string, opts ...string) {
	value := f.Get(field)
	if value == "" {
		return
	}
	for _, opt := range opts {
		if value == opt {
			return
		}
	}
	f.Errors.Add(field, "This field is invalid")
}

func (f *Form) MatchesPattern(field string, pattern *regexp.Regexp) {
	value := f.Get(field)
	if value == "" {
		return
	}
	if !pattern.MatchString(value) {
		f.Errors.Add(field, "This field is invalid")
	}
}

func (f *Form) MatchesOtherField(field string, other string) {
	value := f.Get(field)
	otherValue := f.Get(other)
	if value == "" || otherValue == "" {
		return
	}
	if value != otherValue {
		f.Errors.Add(field, fmt.Sprintf("This field must match the \"%s\" field", other))
	}
}

func (f *Form) Valid() bool {
	return len(f.Errors) == 0
}
