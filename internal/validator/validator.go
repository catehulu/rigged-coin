package validator

import (
	"net/http"
	"net/url"
	"strings"
)

// Form creates a custom form struct and embeds a url.Values object
type Validator struct {
	url.Values
	Errors errors
}

// Valid returns true if there are no errors, otherwise false
func (f *Validator) Valid() bool {
	return len(f.Errors) == 0
}

// New initializes a form struct
func New(data url.Values) *Validator {
	return &Validator{
		data,
		errors(map[string][]string{}),
	}
}

// Required checks for required fields
func (f *Validator) Required(fields ...string) {
	for _, field := range fields {
		value := f.Get(field)
		if strings.TrimSpace(value) == "" {
			f.Errors.Add(field, "This field cannot be blank")
		}
	}
}

// Has checks if Validator field is in post and not empty
func (f *Validator) Has(field string, r *http.Request) bool {
	x := r.Form.Get(field)
	return x != ""
}
