// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: n13t/idm/v0/idm.proto

package n13t_idm_v0

import (
	"bytes"
	"errors"
	"fmt"
	"net"
	"net/mail"
	"net/url"
	"regexp"
	"strings"
	"time"
	"unicode/utf8"

	"github.com/golang/protobuf/ptypes"
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
	_ = ptypes.DynamicAny{}
)

// define the regex for a UUID once up-front
var _idm_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on CreateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *CreateUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return CreateUserRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	// no validation rules for Password

	return nil
}

// CreateUserRequestValidationError is the validation error returned by
// CreateUserRequest.Validate if the designated constraints aren't met.
type CreateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserRequestValidationError) ErrorName() string {
	return "CreateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserRequestValidationError{}

// Validate checks the field values on CreateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *CreateUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// CreateUserResponseValidationError is the validation error returned by
// CreateUserResponse.Validate if the designated constraints aren't met.
type CreateUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e CreateUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e CreateUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e CreateUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e CreateUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e CreateUserResponseValidationError) ErrorName() string {
	return "CreateUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e CreateUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sCreateUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = CreateUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = CreateUserResponseValidationError{}

// Validate checks the field values on UpdateUserRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *UpdateUserRequest) Validate() error {
	if m == nil {
		return nil
	}

	if v, ok := interface{}(m.GetUser()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UpdateUserRequestValidationError{
				field:  "User",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	return nil
}

// UpdateUserRequestValidationError is the validation error returned by
// UpdateUserRequest.Validate if the designated constraints aren't met.
type UpdateUserRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserRequestValidationError) ErrorName() string {
	return "UpdateUserRequestValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserRequestValidationError{}

// Validate checks the field values on UpdateUserResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *UpdateUserResponse) Validate() error {
	if m == nil {
		return nil
	}

	return nil
}

// UpdateUserResponseValidationError is the validation error returned by
// UpdateUserResponse.Validate if the designated constraints aren't met.
type UpdateUserResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UpdateUserResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UpdateUserResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UpdateUserResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UpdateUserResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UpdateUserResponseValidationError) ErrorName() string {
	return "UpdateUserResponseValidationError"
}

// Error satisfies the builtin error interface
func (e UpdateUserResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUpdateUserResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UpdateUserResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UpdateUserResponseValidationError{}

// Validate checks the field values on HealthCheckRequest with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheckRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Service

	return nil
}

// HealthCheckRequestValidationError is the validation error returned by
// HealthCheckRequest.Validate if the designated constraints aren't met.
type HealthCheckRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckRequestValidationError) ErrorName() string {
	return "HealthCheckRequestValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckRequestValidationError{}

// Validate checks the field values on HealthCheckResponse with the rules
// defined in the proto definition for this message. If any rules are
// violated, an error is returned.
func (m *HealthCheckResponse) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Status

	return nil
}

// HealthCheckResponseValidationError is the validation error returned by
// HealthCheckResponse.Validate if the designated constraints aren't met.
type HealthCheckResponseValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e HealthCheckResponseValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e HealthCheckResponseValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e HealthCheckResponseValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e HealthCheckResponseValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e HealthCheckResponseValidationError) ErrorName() string {
	return "HealthCheckResponseValidationError"
}

// Error satisfies the builtin error interface
func (e HealthCheckResponseValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sHealthCheckResponse.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = HealthCheckResponseValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = HealthCheckResponseValidationError{}

// Validate checks the field values on ListUsersRequest with the rules defined
// in the proto definition for this message. If any rules are violated, an
// error is returned.
func (m *ListUsersRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Search

	return nil
}

// ListUsersRequestValidationError is the validation error returned by
// ListUsersRequest.Validate if the designated constraints aren't met.
type ListUsersRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ListUsersRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ListUsersRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ListUsersRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ListUsersRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ListUsersRequestValidationError) ErrorName() string { return "ListUsersRequestValidationError" }

// Error satisfies the builtin error interface
func (e ListUsersRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sListUsersRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ListUsersRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ListUsersRequestValidationError{}

// Validate checks the field values on SumRequest with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SumRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for A

	// no validation rules for B

	return nil
}

// SumRequestValidationError is the validation error returned by
// SumRequest.Validate if the designated constraints aren't met.
type SumRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SumRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SumRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SumRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SumRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SumRequestValidationError) ErrorName() string { return "SumRequestValidationError" }

// Error satisfies the builtin error interface
func (e SumRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSumRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SumRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SumRequestValidationError{}

// Validate checks the field values on SumReply with the rules defined in the
// proto definition for this message. If any rules are violated, an error is returned.
func (m *SumReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for V

	// no validation rules for Err

	return nil
}

// SumReplyValidationError is the validation error returned by
// SumReply.Validate if the designated constraints aren't met.
type SumReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SumReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SumReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SumReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SumReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SumReplyValidationError) ErrorName() string { return "SumReplyValidationError" }

// Error satisfies the builtin error interface
func (e SumReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSumReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SumReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SumReplyValidationError{}

// Validate checks the field values on ConcatRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ConcatRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for A

	// no validation rules for B

	return nil
}

// ConcatRequestValidationError is the validation error returned by
// ConcatRequest.Validate if the designated constraints aren't met.
type ConcatRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcatRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcatRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcatRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcatRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcatRequestValidationError) ErrorName() string { return "ConcatRequestValidationError" }

// Error satisfies the builtin error interface
func (e ConcatRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcatRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcatRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcatRequestValidationError{}

// Validate checks the field values on ConcatReply with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *ConcatReply) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for V

	// no validation rules for Err

	return nil
}

// ConcatReplyValidationError is the validation error returned by
// ConcatReply.Validate if the designated constraints aren't met.
type ConcatReplyValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e ConcatReplyValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e ConcatReplyValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e ConcatReplyValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e ConcatReplyValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e ConcatReplyValidationError) ErrorName() string { return "ConcatReplyValidationError" }

// Error satisfies the builtin error interface
func (e ConcatReplyValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sConcatReply.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = ConcatReplyValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = ConcatReplyValidationError{}
