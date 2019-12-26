// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: idm.proto

package n13t_idm_v1alpha1

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

// Validate checks the field values on SearchRequest with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *SearchRequest) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Query

	// no validation rules for PageNumber

	// no validation rules for ResultPerPage

	return nil
}

// SearchRequestValidationError is the validation error returned by
// SearchRequest.Validate if the designated constraints aren't met.
type SearchRequestValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e SearchRequestValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e SearchRequestValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e SearchRequestValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e SearchRequestValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e SearchRequestValidationError) ErrorName() string { return "SearchRequestValidationError" }

// Error satisfies the builtin error interface
func (e SearchRequestValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sSearchRequest.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = SearchRequestValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = SearchRequestValidationError{}