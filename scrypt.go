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
	"golang.org/x/crypto/scrypt"
	"log"
)

// ScryptConf contains all parameters for scrypt.
// See https://godoc.org/golang.org/x/crypto/scrypt
// This implementation is created with a rounds parameter in contrast to the N parameter in the scrypt library.
// N is just 2**rounds.
// To change this parameter use SetRounds.
type ScryptConf struct {
	n, rounds, R, P, KeyLen int
}

// NewScryptConf returns a new scrypt config.
func NewScryptConf(rounds, r, p, keyLen int) *ScryptConf {
	res := &ScryptConf{
		R:      r,
		P:      p,
		KeyLen: keyLen,
	}
	res.SetRounds(rounds)
	return res
}

// Copy returns a copy of a config.
func (conf *ScryptConf) Copy() *ScryptConf {
	return &ScryptConf{
		rounds: conf.rounds,
		n:      conf.n,
		R:      conf.R,
		P:      conf.P,
		KeyLen: conf.KeyLen,
	}
}

// String returns a human-readable string reprensetation.
func (conf *ScryptConf) String() string {
	return fmt.Sprintf("&{rounds: %d, R: %d, P: %d, KeyLen: %d}", conf.rounds, conf.R, conf.P, conf.KeyLen)
}

// GetN returns the N parameter for scrypt, that is 2 ** rounds.
func (conf *ScryptConf) GetN() int {
	return conf.n
}

// GetRounds returns the rounds parameter (N is 2 ** rounds).
func (conf *ScryptConf) GetRounds() int {
	return conf.rounds
}

// SetRounds sets the rounds parameter and thus N = 2 ** rounds.
func (conf *ScryptConf) SetRounds(rounds int) {
	pow := Pow(2, int64(rounds))
	asInt := int(pow)
	if asInt <= 0 {
		log.Printf("Invalid rounds parameter for scrypt: %d. Using default (16)\n", rounds)
		asInt = 65536
		rounds = 16
	}
	conf.rounds = rounds
	conf.n = asInt
}

// DefaultScryptConf is the default configuration for scrypt.
var DefaultScryptConf = NewScryptConf(16, 8, 1, 64)

// ScryptData stores in addition to a config also the salt and key, both
// base64 encoded (Salt and Key) as well as the raw version (decoded from
// Salt and Key).
// This is the type parsed from an scrypt config.
type ScryptData struct {
	*ScryptConf
	// encoded with base64
	Salt, Key       string
	RawSalt, RawKey []byte
}

// ScryptHasher is a Hasher using scrypt.
type ScryptHasher struct {
	*ScryptConf
}

// NewScryptHasher returns a new ScryptHasher with the given parameters.
func NewScryptHasher(conf *ScryptConf) *ScryptHasher {
	if conf == nil {
		conf = DefaultScryptConf
	}
	return &ScryptHasher{conf}
}

// Copy returns a copy of the hasher.
func (h *ScryptHasher) Copy() *ScryptHasher {
	return &ScryptHasher{h.ScryptConf.Copy()}
}

// key returns the scrypt key of the password given the clear text password,
// salt and parameters from the config of the Hasher.
func (h *ScryptHasher) key(password string, salt []byte) ([]byte, error) {
	return scrypt.Key([]byte(password), salt, h.N, h.R, h.P, h.KeyLen)
}
