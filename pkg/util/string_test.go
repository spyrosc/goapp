package util

import (
	"regexp"
	"testing"
)

func TestRandStringIsHex(t *testing.T) {

	hexRegex := regexp.MustCompile(`^[0-9A-F]+$`)
	numTests := 10000

	for i := 0; i < numTests; i++ {
		randomString := RandString(10)
		if !hexRegex.MatchString(randomString) {
			t.Errorf("RandString() did not return a hex value: %s", randomString)
		}
	}
}

func BenchmarkRandString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		_ = RandString(10)
	}
}
