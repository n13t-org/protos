// Code generated by protoc-gen-validate. DO NOT EDIT.
// source: n13t/idm/v0/user.proto

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
var _user_uuidPattern = regexp.MustCompile("^[0-9a-fA-F]{8}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{4}-[0-9a-fA-F]{12}$")

// Validate checks the field values on User with the rules defined in the proto
// definition for this message. If any rules are violated, an error is returned.
func (m *User) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Enabled

	for key, val := range m.GetAttributes() {
		_ = val

		// no validation rules for Attributes[key]

		if v, ok := interface{}(val).(interface{ Validate() error }); ok {
			if err := v.Validate(); err != nil {
				return UserValidationError{
					field:  fmt.Sprintf("Attributes[%v]", key),
					reason: "embedded message failed validation",
					cause:  err,
				}
			}
		}

	}

	if err := m._validateUuid(m.GetSub()); err != nil {
		return UserValidationError{
			field:  "Sub",
			reason: "value must be a valid UUID",
			cause:  err,
		}
	}

	if len(m.GetName()) > 256 {
		return UserValidationError{
			field:  "Name",
			reason: "value length must be at most 256 bytes",
		}
	}

	if !_User_Name_Pattern.MatchString(m.GetName()) {
		return UserValidationError{
			field:  "Name",
			reason: "value does not match regex pattern \"^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$\"",
		}
	}

	if len(m.GetGivenName()) > 64 {
		return UserValidationError{
			field:  "GivenName",
			reason: "value length must be at most 64 bytes",
		}
	}

	if !_User_GivenName_Pattern.MatchString(m.GetGivenName()) {
		return UserValidationError{
			field:  "GivenName",
			reason: "value does not match regex pattern \"^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$\"",
		}
	}

	if len(m.GetFamilyName()) > 64 {
		return UserValidationError{
			field:  "FamilyName",
			reason: "value length must be at most 64 bytes",
		}
	}

	if !_User_FamilyName_Pattern.MatchString(m.GetFamilyName()) {
		return UserValidationError{
			field:  "FamilyName",
			reason: "value does not match regex pattern \"^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$\"",
		}
	}

	if len(m.GetMiddleName()) > 64 {
		return UserValidationError{
			field:  "MiddleName",
			reason: "value length must be at most 64 bytes",
		}
	}

	if !_User_MiddleName_Pattern.MatchString(m.GetMiddleName()) {
		return UserValidationError{
			field:  "MiddleName",
			reason: "value does not match regex pattern \"^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$\"",
		}
	}

	if len(m.GetNickname()) > 256 {
		return UserValidationError{
			field:  "Nickname",
			reason: "value length must be at most 256 bytes",
		}
	}

	if !_User_Nickname_Pattern.MatchString(m.GetNickname()) {
		return UserValidationError{
			field:  "Nickname",
			reason: "value does not match regex pattern \"^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$\"",
		}
	}

	// no validation rules for PreferredUsername

	if uri, err := url.Parse(m.GetProfile()); err != nil {
		return UserValidationError{
			field:  "Profile",
			reason: "value must be a valid URI",
			cause:  err,
		}
	} else if !uri.IsAbs() {
		return UserValidationError{
			field:  "Profile",
			reason: "value must be absolute",
		}
	}

	if uri, err := url.Parse(m.GetPicture()); err != nil {
		return UserValidationError{
			field:  "Picture",
			reason: "value must be a valid URI",
			cause:  err,
		}
	} else if !uri.IsAbs() {
		return UserValidationError{
			field:  "Picture",
			reason: "value must be absolute",
		}
	}

	if uri, err := url.Parse(m.GetWebsite()); err != nil {
		return UserValidationError{
			field:  "Website",
			reason: "value must be a valid URI",
			cause:  err,
		}
	} else if !uri.IsAbs() {
		return UserValidationError{
			field:  "Website",
			reason: "value must be absolute",
		}
	}

	if !_User_Email_Pattern.MatchString(m.GetEmail()) {
		return UserValidationError{
			field:  "Email",
			reason: "value does not match regex pattern \"^\\\\w+([-+.']\\\\w+)*@\\\\w+([-.]\\\\w+)*\\\\.\\\\w+([-.]\\\\w+)*$\"",
		}
	}

	// no validation rules for EmailVerified

	// no validation rules for Gender

	// no validation rules for Birthdate

	if !_User_Zoneinfo_Pattern.MatchString(m.GetZoneinfo()) {
		return UserValidationError{
			field:  "Zoneinfo",
			reason: "value does not match regex pattern \"^\\\\w+(/\\\\w+){1,2}([-|+][0-9]{1,2})?$\"",
		}
	}

	// no validation rules for Locale

	// no validation rules for PhoneNumber

	// no validation rules for PhoneNumberVerified

	if v, ok := interface{}(m.GetAddress()).(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return UserValidationError{
				field:  "Address",
				reason: "embedded message failed validation",
				cause:  err,
			}
		}
	}

	if m.GetUpdatedAt() == nil {
		return UserValidationError{
			field:  "UpdatedAt",
			reason: "value is required",
		}
	}

	// no validation rules for Username

	return nil
}

func (m *User) _validateUuid(uuid string) error {
	if matched := _user_uuidPattern.MatchString(uuid); !matched {
		return errors.New("invalid uuid format")
	}

	return nil
}

// UserValidationError is the validation error returned by User.Validate if the
// designated constraints aren't met.
type UserValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e UserValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e UserValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e UserValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e UserValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e UserValidationError) ErrorName() string { return "UserValidationError" }

// Error satisfies the builtin error interface
func (e UserValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = UserValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = UserValidationError{}

var _User_Name_Pattern = regexp.MustCompile("^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$")

var _User_GivenName_Pattern = regexp.MustCompile("^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$")

var _User_FamilyName_Pattern = regexp.MustCompile("^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$")

var _User_MiddleName_Pattern = regexp.MustCompile("^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$")

var _User_Nickname_Pattern = regexp.MustCompile("^[^[0-9]A-Za-z]+( [^[0-9]A-Za-z]+)*$")

var _User_Email_Pattern = regexp.MustCompile("^\\w+([-+.']\\w+)*@\\w+([-.]\\w+)*\\.\\w+([-.]\\w+)*$")

var _User_Zoneinfo_Pattern = regexp.MustCompile("^\\w+(/\\w+){1,2}([-|+][0-9]{1,2})?$")

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

	// no validation rules for Err

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

	// no validation rules for Err

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

// Validate checks the field values on User_Address with the rules defined in
// the proto definition for this message. If any rules are violated, an error
// is returned.
func (m *User_Address) Validate() error {
	if m == nil {
		return nil
	}

	// no validation rules for Formatted

	// no validation rules for StreetAddress

	// no validation rules for Locality

	// no validation rules for Region

	// no validation rules for PostalCode

	// no validation rules for Country

	return nil
}

// User_AddressValidationError is the validation error returned by
// User_Address.Validate if the designated constraints aren't met.
type User_AddressValidationError struct {
	field  string
	reason string
	cause  error
	key    bool
}

// Field function returns field value.
func (e User_AddressValidationError) Field() string { return e.field }

// Reason function returns reason value.
func (e User_AddressValidationError) Reason() string { return e.reason }

// Cause function returns cause value.
func (e User_AddressValidationError) Cause() error { return e.cause }

// Key function returns key value.
func (e User_AddressValidationError) Key() bool { return e.key }

// ErrorName returns error name.
func (e User_AddressValidationError) ErrorName() string { return "User_AddressValidationError" }

// Error satisfies the builtin error interface
func (e User_AddressValidationError) Error() string {
	cause := ""
	if e.cause != nil {
		cause = fmt.Sprintf(" | caused by: %v", e.cause)
	}

	key := ""
	if e.key {
		key = "key for "
	}

	return fmt.Sprintf(
		"invalid %sUser_Address.%s: %s%s",
		key,
		e.field,
		e.reason,
		cause)
}

var _ error = User_AddressValidationError{}

var _ interface {
	Field() string
	Reason() string
	Key() bool
	Cause() error
	ErrorName() string
} = User_AddressValidationError{}
