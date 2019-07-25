// +build gofuzz

/*
 * The MIT License (MIT)
 * Copyright (c) 2019 Thibault NORMAND
 *
 * Permission is hereby granted, free of charge, to any person obtaining a copy
 * of this software and associated documentation files (the "Software"), to deal
 * in the Software without restriction, including without limitation the rights
 * to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
 * copies of the Software, and to permit persons to whom the Software is
 * furnished to do so, subject to the following conditions:
 *
 * The above copyright notice and this permission notice shall be included in
 * all copies or substantial portions of the Software.
 *
 * THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
 * IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
 * FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
 * AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
 * LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
 * OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN
 * THE SOFTWARE.
 */

package hasher

// Fuzz usage:
//   go get github.com/dvyukov/go-fuzz/...
//
//   go-fuzz-build go.zenithar.org/butcher && go-fuzz -bin=./butcher-fuzz.zip -workdir=/tmp/butcher-fuzz
func Fuzz(data []byte) int {
	hResult, err := Hash(data)
	if err != nil {
		if hResult != "" {
			panic("hResult != '' on error")
		}
		return 0
	}
	valid, err := Verify([]byte(hResult), data)
	if err != nil {
		if valid {
			panic("valid is true on error")
		}
		return 2
	}
	return 1
}
