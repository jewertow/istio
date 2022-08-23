// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: envoy/extensions/filters/network/connection_limit/v3/connection_limit.proto

package connection_limitv3

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"sort"
	"strings"
	"time"
	"unicode/utf8"

	"google.golang.org/protobuf/types/known/anypb"
)

// ensure the imports are used
var (
	_ = bytes.MinRead
	_ = errors.New("")
	_ = fmt.Print
	_ = utf8.UTFMax
	_ = (*regexp.Regexp)(nil)
	_ = (*strings.Reader)(nil)
	_ = net.IPv4len
	_ = time.Duration(0)
	_ = (*url.URL)(nil)
	_ = (*mail.Address)(nil)
	_ = anypb.Any{}
	_ = sort.Sort
)

// Validate checks the field values on ConnectionLimit with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *ConnectionLimit) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on ConnectionLimit with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// ConnectionLimitMultiError, or nil if none found.
func (m *ConnectionLimit) ValidateAll() error {
	return m.validate(true)
}

func (m *ConnectionLimit) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetStatPrefix()) < 1 {
		err := ConnectionLimitValidationError{
			field:  "StatPrefix",
			reason: "value length must be at least 1 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if wrapper := m.GetMaxConnections(); wrapper != nil {

		if wrapper.GetValue() < 1 {
			err := ConnectionLimitValidationError{
				field:  "MaxConnections",
				reason: "value must be greater than or equal to 1",
			}
			if !all {
				return err
			}
			errors = append(errors, err)
		}

	}

	if all {
		switch v := interface{}(m.GetDelay()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConnectionLimitValidationError{
					field:  "Delay",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConnectionLimitValidationError{
					field:  "Delay",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetDelay()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConnectionLimitValidationError{
				field:  "Delay",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if all {
		switch v := interface{}(m.GetRuntimeEnabled()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, ConnectionLimitValidationError{
					field:  "RuntimeEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, ConnectionLimitValidationError{
					field:  "RuntimeEnabled",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetRuntimeEnabled()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return ConnectionLimitValidationError{
				field:  "RuntimeEnabled",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return ConnectionLimitMultiError(errors)
	}

	return nil
}

// ConnectionLimitMultiError is an error wrapping multiple validation errors
// returned by ConnectionLimit.ValidateAll() if the designated constraints
// aren't met.
type ConnectionLimitMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m ConnectionLimitMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m ConnectionLimitMultiError) AllErrors() []error { return m }

// ConnectionLimitValidationError is the validation error returned by
// ConnectionLimit.Validate if the designated constraints aren't met.
type ConnectionLimitValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConnectionLimitValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConnectionLimitValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConnectionLimitValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConnectionLimitValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConnectionLimitValidationError) ErrorName() string { return "ConnectionLimitValidationError" }

// Error satisfies the builtin error interface
func (e ConnectionLimitValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConnectionLimit.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConnectionLimitValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConnectionLimitValidationError{}
