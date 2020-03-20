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

// Hasher is the general interface for creating hashed versions of passwords.
// The returned slice of bytes contains the hashes as well as all other information, depending on the algorithm.
// For example the algorithm might add a prefix and configuration variables if it needs to.
// For example see https://openwall.info/wiki/john/sample-hashes.
type Hasher interface {
	Generate(password string) ([]byte, error)
}

// Validator is an interface that provides a method to compare the hashed version
// of a password with a given clear text version. Any error returned should
// be considered as an authentication fail. Only a nil return value indicates
// success.
//
// There are some predefined errors that can help you to narrow the cause.
// But not all implementation are required to use these errors.
// Special errors include: SyntaxError if there is an error in the expected syntax.
// VersionError: If the version used to create the hash value is not
// compatible with the implemented algorithm.
// PasswordMismatchError: If the clear text version is not the password used to
// create the hash.
//
// For comparing hashes (as bytes) you should use CompareHashes.
//
// Note that a validator implementation provides validation for a specific
// hashing algorithm, like one implementation for bcrypt, scrypt etc.
type Validator interface {
	Compare(hashed []byte, password string) error
}
