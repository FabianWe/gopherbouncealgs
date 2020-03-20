// Copyright 2019 - 2020 Fabian Wenzelmann <fabianwen@posteo.eu>
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
	"crypto/rand"
	"crypto/subtle"
	"io"
)

// Pow computes base**exp. It only works with positive integers.
func Pow(base, exp int64) int64 {
	var result int64 = 1
	for {
		if (exp & 1) != 0 {
			result *= base
		}
		exp >>= 1
		if exp == 0 {
			break
		}
		base *= base
	}
	return result
}

// GenSalt computes a cryptographically secure salt given the number of bytes.
func GenSalt(numBytes int) ([]byte, error) {
	salt := make([]byte, numBytes)
	_, err := io.ReadFull(rand.Reader, salt)
	return salt, err
}

// CompareHashes uses a constant time compare algorithm to compare to key
// hashes. Constant time compare functions are important or otherwise attackers
// might infer knowledge about the real password.
func CompareHashes(x, y []byte) bool {
	return subtle.ConstantTimeCompare(x, y) == 1
}
