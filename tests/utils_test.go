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

package tests

import (
	"github.com/FabianWe/gopherbouncealgs"
	"testing"
)

// TestPow tests utils.Pow.
func TestPow(t *testing.T) {
	tests := []struct {
		base, exp, out int64
	}{
		{42, 0, 1},
		{0, 42, 0},
		{2, 8, 256},
		{10, 10, 10000000000},
	}
	for _, tc := range tests {
		if got := gopherbouncealgs.Pow(tc.base, tc.exp); got != tc.out {
			t.Errorf("Expected that %d ^ %d = %d, got %d", tc.base, tc.exp, tc.out, got)
		}
	}
}

// TestCompareHashes tests utils.CompareHashes.
func TestCompareHashes(t *testing.T) {
	tests := []struct {
		first, second []byte
		out           bool
	}{
		{[]byte("abcd"), []byte("abcd"), true},
		{[]byte("abcd"), []byte("abc"), false},
	}
	for _, tc := range tests {
		if got := gopherbouncealgs.CompareHashes(tc.first, tc.second); got != tc.out {
			t.Errorf("Expected that (bytes of) %s == %s ==> %v, got %v", string(tc.first),
				string(tc.second), tc.out, got)
		}
	}
}
