package tokenizer

import (
	"json-parser/internal/tokenizer"
	"testing"
)


func TestJsonStringSuccessfullyTokenized(t *testing.T) {
  jsonString := `{"test":"value"}`
  got := tokenizer.Tokenize(jsonString)
  want := []rune{}

  if (len(got) != len(want)) {
    t.Fatalf("got: %+q, want: %+q", got, want)
  }
}
