package basic_test

import (
	"testing"
	"unicode/utf8"

	"github.com/veegrace/go-fuzzing/basic"
)

func FuzzBasicOverwriteString(f *testing.F) {
	f.Add("Hello, world!", 'A', 15)

	f.Fuzz(func(t *testing.T, str string, value rune, count int) {
		basic.OverwriteString(str, value, count)
	})
}

func FuzzOverwriteStringSuffix(f *testing.F) {
	f.Add("Hello, world!", 'A', 15)

	f.Fuzz(func(t *testing.T, str string, value rune, count int) {
		result := basic.OverwriteString(str, value, count)
		if count > 0 && count < utf8.RuneCountInString(str) {
			resultSuffix := string([]rune(result)[count:])
			strSuffix := string([]rune(str)[count:])
			if resultSuffix != strSuffix {
				t.Fatalf("OverwriteString modified too many characters! Expected %s, got %s.", strSuffix, resultSuffix)
			}
		}
	})
}
