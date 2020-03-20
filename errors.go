// Copyright 2018 - 2020 Fabian Wenzelmann <fabianwen@posteo.eu>
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
// http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package gopherbouncealgs

import (
	"fmt"
	"strings"
)

// SyntaxError is used if a validator expected a different syntax.
type SyntaxError string

// NewSyntaxError returns a new SyntaxError.
func NewSyntaxError(cause string) SyntaxError {
	return SyntaxError(cause)
}

// Error returns the string representation of the error.
func (err SyntaxError) Error() string {
	return "Syntax error: " + string(err)
}

// String returns the string representation of the error.
func (err SyntaxError) String() string {
	return err.Error()
}

// VersionError is an error returned if the version of a hashed version is not compatible with the library.
type VersionError struct {
	Prefix   string
	Expected []string
	Got      string
}

// NewVersionError returns a new VersionError.
// Prefix is appended to the error message.
// Expected is a list of allowed versions and got is the version that is not supported.
func NewVersionError(prefix string, expected []string, got string) VersionError {
	return VersionError{
		Prefix:   prefix,
		Expected: expected,
		Got:      got,
	}
}

// Error returns the string representation of the error.
func (err VersionError) Error() string {
	var errorString string
	if len(err.Expected) == 1 {
		errorString = fmt.Sprintf("Invalid algorithm version, expected %s, got %s", err.Expected, err.Got)
	} else {
		errorString = fmt.Sprintf("Invalid algorithm version, expected one of %s, got %s", strings.Join(err.Expected, ", "), err.Got)
	}
	if err.Prefix == "" {
		return "Algorithm error: " + errorString
	}
	return fmt.Sprintf("%s: %s", err.Prefix, errorString)
}

// PasswordMismatchError is returned if a hashed version and a clear text
// version do not match. No password details are made public by this error.
type PasswordMismatchError struct{}

// NewPasswordMismatchError returns a new PasswordMismatchError.
func NewPasswordMismatchError() PasswordMismatchError {
	return PasswordMismatchError{}
}

// Error returns the string representation of the error.
func (err PasswordMismatchError) Error() string {
	return "gopherbouncealg: Passwords do not match"
}
