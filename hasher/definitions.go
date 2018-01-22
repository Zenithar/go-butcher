/*
 * The MIT License (MIT)
 * Copyright (c) 2018 Thibault NORMAND
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy of this software
 * and associated documentation files (the "Software"), to deal in the Software without restriction,
 * including without limitation the rights to use, copy, modify, merge, publish, distribute, sublicense,
 * and/or sell copies of the Software, and to permit persons to whom the Software is furnished to do so,
 * subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in all copies or substantial
 * portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED
 * TO THE WARRANTIES OF MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL
 * THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
 * TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
 */

package hasher

import (
	"crypto/sha512"
	"hash"

	"golang.org/x/crypto/blake2b"
	"golang.org/x/crypto/sha3"
)

const (
	// Argon2i defines the argon2i hashing algorithm
	Argon2i = "argon2i"
	//ScryptBlake2b512 defines scrypt+blake2b-512 hashing algorithm
	ScryptBlake2b512 = "scrypt+blake2b-512"
	// Pbkdf2Blake2b512 defines pbkdf2+blake2b-512 hashing algorithm
	Pbkdf2Blake2b512 = "pbkdf2+blake2b-512"
	// Pbkdf2Sha512 defines pbkdf2+sha512 hashing algorithm
	Pbkdf2Sha512 = "pbkdf2+sha512"
	// Pbkdf2Keccak512 defines pbkdf2+sha3-512 hashing algorithm
	Pbkdf2Keccak512 = "pbkdf2+sha3-512"
)

// Strategies defines available hashing strategies
var Strategies = map[string]func([]byte) Strategy{
	Argon2i: func(salt []byte) Strategy {
		s, _ := newArgon2Deriver(salt)
		return s
	},
	ScryptBlake2b512: func(salt []byte) Strategy {
		s, _ := newScryptDeriver(func() hash.Hash {
			h, _ := blake2b.New512(nil)
			return h
		}, salt, 12)
		return s
	},
	Pbkdf2Blake2b512: func(salt []byte) Strategy {
		s, _ := newPbkdf2Deriver(func() hash.Hash {
			h, _ := blake2b.New512(nil)
			return h
		}, salt, 50000, blake2b.Size)
		return s
	},
	Pbkdf2Sha512: func(salt []byte) Strategy {
		s, _ := newPbkdf2Deriver(sha512.New, salt, 50000, sha512.Size)
		return s
	},
	Pbkdf2Keccak512: func(salt []byte) Strategy {
		s, _ := newPbkdf2Deriver(sha3.New512, salt, 50000, 64)
		return s
	},
}
