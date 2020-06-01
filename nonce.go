// Licensed to Butcher under one or more contributor
// license agreements. See the NOTICE file distributed with
// this work for additional information regarding copyright
// ownership. Butcher licenses this file to you under
// the Apache License, Version 2.0 (the "License"); you may
// not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing,
// software distributed under the License is distributed on an
// "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
// KIND, either express or implied.  See the License for the
// specific language governing permissions and limitations
// under the License.

package butcher

import "crypto/rand"

// FixedNonce returns a nonce factory that returns the given salt
func FixedNonce(salt []byte) func() []byte {
	return func() []byte {
		return salt
	}
}

// RandomNonce returns a nonce factory that returns a random length bound salt
func RandomNonce(length int) func() []byte {
	return func() []byte {
		salt := make([]byte, length)
		if _, err := rand.Read(salt); err != nil {
			panic(err)
		}
		return salt
	}
}
