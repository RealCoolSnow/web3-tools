// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: api/shop/service/v1/shop.proto

package v1

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

// Validate checks the field values on RegisterRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *RegisterRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// RegisterRequestMultiError, or nil if none found.
func (m *RegisterRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMobile()) != 11 {
		err := RegisterRequestValidationError{
			field:  "Mobile",
			reason: "value length must be 11 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if l := utf8.RuneCountInString(m.GetPass()); l < 6 || l > 18 {
		err := RegisterRequestValidationError{
			field:  "Pass",
			reason: "value length must be between 6 and 18 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	// no validation rules for NickName

	// no validation rules for Age

	if len(errors) > 0 {
		return RegisterRequestMultiError(errors)
	}

	return nil
}

// RegisterRequestMultiError is an error wrapping multiple validation errors
// returned by RegisterRequest.ValidateAll() if the designated constraints
// aren't met.
type RegisterRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterRequestMultiError) AllErrors() []error { return m }

// RegisterRequestValidationError is the validation error returned by
// RegisterRequest.Validate if the designated constraints aren't met.
type RegisterRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterRequestValidationError) ErrorName() string { return "RegisterRequestValidationError" }

// Error satisfies the builtin error interface
func (e RegisterRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterRequestValidationError{}

// Validate checks the field values on RegisterReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *RegisterReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on RegisterReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in RegisterReplyMultiError, or
// nil if none found.
func (m *RegisterReply) ValidateAll() error {
	return m.validate(true)
}

func (m *RegisterReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for Mobile

	// no validation rules for NickName

	// no validation rules for Age

	if len(errors) > 0 {
		return RegisterReplyMultiError(errors)
	}

	return nil
}

// RegisterReplyMultiError is an error wrapping multiple validation errors
// returned by RegisterReply.ValidateAll() if the designated constraints
// aren't met.
type RegisterReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m RegisterReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m RegisterReplyMultiError) AllErrors() []error { return m }

// RegisterReplyValidationError is the validation error returned by
// RegisterReply.Validate if the designated constraints aren't met.
type RegisterReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e RegisterReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e RegisterReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e RegisterReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e RegisterReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e RegisterReplyValidationError) ErrorName() string { return "RegisterReplyValidationError" }

// Error satisfies the builtin error interface
func (e RegisterReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sRegisterReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = RegisterReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = RegisterReplyValidationError{}

// Validate checks the field values on LoginRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginRequestMultiError, or
// nil if none found.
func (m *LoginRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if utf8.RuneCountInString(m.GetMobile()) != 11 {
		err := LoginRequestValidationError{
			field:  "Mobile",
			reason: "value length must be 11 runes",
		}
		if !all {
			return err
		}
		errors = append(errors, err)

	}

	if l := utf8.RuneCountInString(m.GetPass()); l < 6 || l > 18 {
		err := LoginRequestValidationError{
			field:  "Pass",
			reason: "value length must be between 6 and 18 runes, inclusive",
		}
		if !all {
			return err
		}
		errors = append(errors, err)
	}

	if len(errors) > 0 {
		return LoginRequestMultiError(errors)
	}

	return nil
}

// LoginRequestMultiError is an error wrapping multiple validation errors
// returned by LoginRequest.ValidateAll() if the designated constraints aren't met.
type LoginRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginRequestMultiError) AllErrors() []error { return m }

// LoginRequestValidationError is the validation error returned by
// LoginRequest.Validate if the designated constraints aren't met.
type LoginRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginRequestValidationError) ErrorName() string { return "LoginRequestValidationError" }

// Error satisfies the builtin error interface
func (e LoginRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginRequestValidationError{}

// Validate checks the field values on LoginReply with the rules defined in the
// proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *LoginReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in LoginReplyMultiError, or
// nil if none found.
func (m *LoginReply) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, LoginReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, LoginReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return LoginReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Msg

	if len(errors) > 0 {
		return LoginReplyMultiError(errors)
	}

	return nil
}

// LoginReplyMultiError is an error wrapping multiple validation errors
// returned by LoginReply.ValidateAll() if the designated constraints aren't met.
type LoginReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReplyMultiError) AllErrors() []error { return m }

// LoginReplyValidationError is the validation error returned by
// LoginReply.Validate if the designated constraints aren't met.
type LoginReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReplyValidationError) ErrorName() string { return "LoginReplyValidationError" }

// Error satisfies the builtin error interface
func (e LoginReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReplyValidationError{}

// Validate checks the field values on GetUserRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserRequestMultiError,
// or nil if none found.
func (m *GetUserRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	if len(errors) > 0 {
		return GetUserRequestMultiError(errors)
	}

	return nil
}

// GetUserRequestMultiError is an error wrapping multiple validation errors
// returned by GetUserRequest.ValidateAll() if the designated constraints
// aren't met.
type GetUserRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserRequestMultiError) AllErrors() []error { return m }

// GetUserRequestValidationError is the validation error returned by
// GetUserRequest.Validate if the designated constraints aren't met.
type GetUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserRequestValidationError) ErrorName() string { return "GetUserRequestValidationError" }

// Error satisfies the builtin error interface
func (e GetUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserRequestValidationError{}

// Validate checks the field values on GetUserReply with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *GetUserReply) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserReply with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in GetUserReplyMultiError, or
// nil if none found.
func (m *GetUserReply) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserReply) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Code

	if all {
		switch v := interface{}(m.GetData()).(type) {
		case interface{ ValidateAll() error }:
			if err := v.ValidateAll(); err != nil {
				errors = append(errors, GetUserReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		case interface{ Validate() error }:
			if err := v.Validate(); err != nil {
				errors = append(errors, GetUserReplyValidationError{
					field:  "Data",
					reason: "embedded message failed validation",
					cause:  err,
				})
			}
		}
	} else if v, ok := interface{}(m.GetData()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return GetUserReplyValidationError{
				field:  "Data",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if len(errors) > 0 {
		return GetUserReplyMultiError(errors)
	}

	return nil
}

// GetUserReplyMultiError is an error wrapping multiple validation errors
// returned by GetUserReply.ValidateAll() if the designated constraints aren't met.
type GetUserReplyMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserReplyMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserReplyMultiError) AllErrors() []error { return m }

// GetUserReplyValidationError is the validation error returned by
// GetUserReply.Validate if the designated constraints aren't met.
type GetUserReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReplyValidationError) ErrorName() string { return "GetUserReplyValidationError" }

// Error satisfies the builtin error interface
func (e GetUserReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReplyValidationError{}

// Validate checks the field values on DemoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DemoRequest) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DemoRequest with the rules defined in
// the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DemoRequestMultiError, or
// nil if none found.
func (m *DemoRequest) ValidateAll() error {
	return m.validate(true)
}

func (m *DemoRequest) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for NickName

	if len(errors) > 0 {
		return DemoRequestMultiError(errors)
	}

	return nil
}

// DemoRequestMultiError is an error wrapping multiple validation errors
// returned by DemoRequest.ValidateAll() if the designated constraints aren't met.
type DemoRequestMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DemoRequestMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DemoRequestMultiError) AllErrors() []error { return m }

// DemoRequestValidationError is the validation error returned by
// DemoRequest.Validate if the designated constraints aren't met.
type DemoRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DemoRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DemoRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DemoRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DemoRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DemoRequestValidationError) ErrorName() string { return "DemoRequestValidationError" }

// Error satisfies the builtin error interface
func (e DemoRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDemoRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DemoRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DemoRequestValidationError{}

// Validate checks the field values on DemoResponse with the rules defined in
// the proto definition for this message. If any rules are violated, the first
// error encountered is returned, or nil if there are no violations.
func (m *DemoResponse) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on DemoResponse with the rules defined
// in the proto definition for this message. If any rules are violated, the
// result is a list of violation errors wrapped in DemoResponseMultiError, or
// nil if none found.
func (m *DemoResponse) ValidateAll() error {
	return m.validate(true)
}

func (m *DemoResponse) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Id

	// no validation rules for NickName

	if len(errors) > 0 {
		return DemoResponseMultiError(errors)
	}

	return nil
}

// DemoResponseMultiError is an error wrapping multiple validation errors
// returned by DemoResponse.ValidateAll() if the designated constraints aren't met.
type DemoResponseMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m DemoResponseMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m DemoResponseMultiError) AllErrors() []error { return m }

// DemoResponseValidationError is the validation error returned by
// DemoResponse.Validate if the designated constraints aren't met.
type DemoResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e DemoResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e DemoResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e DemoResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e DemoResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e DemoResponseValidationError) ErrorName() string { return "DemoResponseValidationError" }

// Error satisfies the builtin error interface
func (e DemoResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sDemoResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = DemoResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = DemoResponseValidationError{}

// Validate checks the field values on LoginReply_Data with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *LoginReply_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on LoginReply_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// LoginReply_DataMultiError, or nil if none found.
func (m *LoginReply_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *LoginReply_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Token

	if len(errors) > 0 {
		return LoginReply_DataMultiError(errors)
	}

	return nil
}

// LoginReply_DataMultiError is an error wrapping multiple validation errors
// returned by LoginReply_Data.ValidateAll() if the designated constraints
// aren't met.
type LoginReply_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m LoginReply_DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m LoginReply_DataMultiError) AllErrors() []error { return m }

// LoginReply_DataValidationError is the validation error returned by
// LoginReply_Data.Validate if the designated constraints aren't met.
type LoginReply_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e LoginReply_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e LoginReply_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e LoginReply_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e LoginReply_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e LoginReply_DataValidationError) ErrorName() string { return "LoginReply_DataValidationError" }

// Error satisfies the builtin error interface
func (e LoginReply_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sLoginReply_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = LoginReply_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = LoginReply_DataValidationError{}

// Validate checks the field values on GetUserReply_Data with the rules defined
// in the proto definition for this message. If any rules are violated, the
// first error encountered is returned, or nil if there are no violations.
func (m *GetUserReply_Data) Validate() error {
	return m.validate(false)
}

// ValidateAll checks the field values on GetUserReply_Data with the rules
// defined in the proto definition for this message. If any rules are
// violated, the result is a list of violation errors wrapped in
// GetUserReply_DataMultiError, or nil if none found.
func (m *GetUserReply_Data) ValidateAll() error {
	return m.validate(true)
}

func (m *GetUserReply_Data) validate(all bool) error {
	if m == nil {
		return nil
	}

	var errors []error

	// no validation rules for Name

	if len(errors) > 0 {
		return GetUserReply_DataMultiError(errors)
	}

	return nil
}

// GetUserReply_DataMultiError is an error wrapping multiple validation errors
// returned by GetUserReply_Data.ValidateAll() if the designated constraints
// aren't met.
type GetUserReply_DataMultiError []error

// Error returns a concatenation of all the error messages it wraps.
func (m GetUserReply_DataMultiError) Error() string {
	var msgs []string
	for _, err := range m {
		msgs = append(msgs, err.Error())
	}
	return strings.Join(msgs, "; ")
}

// AllErrors returns a list of validation violation errors.
func (m GetUserReply_DataMultiError) AllErrors() []error { return m }

// GetUserReply_DataValidationError is the validation error returned by
// GetUserReply_Data.Validate if the designated constraints aren't met.
type GetUserReply_DataValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e GetUserReply_DataValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e GetUserReply_DataValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e GetUserReply_DataValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e GetUserReply_DataValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e GetUserReply_DataValidationError) ErrorName() string {
	return "GetUserReply_DataValidationError"
}

// Error satisfies the builtin error interface
func (e GetUserReply_DataValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sGetUserReply_Data.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = GetUserReply_DataValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = GetUserReply_DataValidationError{}
