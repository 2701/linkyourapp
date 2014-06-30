// Example usage:
//
//	s := idgen.New() // s is now "apHCJBl7L1OmC57n"
//
// A standard string created by New() is 16 bytes in length and consists of
// Latin upper and lowercase letters, and numbers (from the set of 62 allowed
// characters), which means that it has ~95 bits of entropy. To get more
// entropy, you can use NewLen(UUIDLen), which returns 20-byte string, giving
// ~119 bits of entropy, or any other desired length.
//
// Functions read from crypto/rand random source, and panic if they fail to
// read from it.
//
// This file is basically a ripoff of https://github.com/dchest/uniuri
// but using math/rand instead of crypt/rand
package lya

import (
	"math/rand"
	"time"
)

const (
	// Standard length of idgen string to achive ~95 bits of entropy.
	StdLen = 16
	// Length of uniurl string to achive ~119 bits of entropy, closest
	// to what can be losslessly converted to UUIDv4 (122 bits).
	UUIDLen = 20
)

func init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

// Standard characters allowed in idgen string.
var StdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789_-")

// New returns a new random string of the standard length, consisting of
// standard characters.
func NewId() string {
	return NewLenChars(StdLen, StdChars)
}

// NewLenChars returns a new random string of the provided length, consisting
// of the provided byte slice of allowed characters.
func NewLenChars(length int, chars []byte) string {
	bytes := make([]byte, length)
	maxlen := len(chars)
	for i := range bytes {
		bytes[i] = chars[rand.Intn(int(maxlen))]
	}
	return string(bytes)
}
