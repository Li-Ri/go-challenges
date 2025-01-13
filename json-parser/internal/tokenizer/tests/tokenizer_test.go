package tokenizer

import (
	"fmt"
	"json-parser/internal/tokenizer"
	"testing"
)


func TestJsonStringSuccessfullyTokenizedWithKeyString(t *testing.T) {
  jsonString := `{"test":"value"}`
  got := tokenizer.Tokenize(jsonString)
  
	want := []tokenizer.Token{ 
		{ Name: tokenizer.BraceOpen, Value: "{" },
		{ Name: tokenizer.DoubleQuotes, Value: "test" },
		{ Name: tokenizer.Colon, Value: ":" },
		{ Name: tokenizer.DoubleQuotes, Value: "value" },
		{ Name: tokenizer.BraceClose, Value: "}" },
	}

	assertTokenSlicesAreEqual(t, got, want)
}

func TestJsonStringSuccessfullyTokenizedWithKeyNumber(t *testing.T) {
jsonString := `{"test":123}`
	got := tokenizer.Tokenize(jsonString)
	want := []tokenizer.Token{
		{ Name: tokenizer.BraceOpen, Value: "{" },
		{ Name: tokenizer.DoubleQuotes, Value: "test" },
		{ Name: tokenizer.Colon, Value: ":" },
		{ Name: tokenizer.Number, Value: "123" },
		{ Name: tokenizer.BraceClose, Value: "}" },
	}
	assertTokenSlicesAreEqual(t, got, want)
}

func TestJsonStringSuccessfullyTokenizedWithBooleanValue(t *testing.T) {
	jsonString := `{"testing": true}`
	got := tokenizer.Tokenize(jsonString)
	want := []tokenizer.Token{
		{ Name: tokenizer.BraceOpen, Value: "{" },
		{ Name: tokenizer.DoubleQuotes, Value: "testing" },
		{ Name: tokenizer.Colon, Value: ":" },
		{ Name: tokenizer.True, Value: "true" },
		{ Name: tokenizer.BraceClose, Value: "}" },
	}
	assertTokenSlicesAreEqual(t, got, want)
}


func TestJsonStringSuccessfullyTokenizedWithArray(t *testing.T) {
	jsonString := `{"test": [1,2,3]}`
	got := tokenizer.Tokenize(jsonString)
	want := []tokenizer.Token{
		{ Name: tokenizer.BraceOpen, Value: "{" },
		{ Name: tokenizer.DoubleQuotes, Value: "test" },
		{ Name: tokenizer.Colon, Value: ":" },
		{ Name: tokenizer.BracketOpen, Value: "[" },
		{ Name: tokenizer.Number, Value: "1" },
		{ Name: tokenizer.Comma, Value: "," },
		{ Name: tokenizer.Number, Value: "2" },
		{ Name: tokenizer.Comma, Value: "," },
		{ Name: tokenizer.Number, Value: "3" },
		{ Name: tokenizer.BracketClose, Value: "]" },
		{ Name: tokenizer.BraceClose, Value: "}" },
	}
	assertTokenSlicesAreEqual(t, got, want)
}

func TestJsonStringSuccessfullyTokenizedWithNestedObjectWithString(t *testing.T) {
	jsonString := `{"test": {"nested": "value"}}`
	got := tokenizer.Tokenize(jsonString)
	want := []tokenizer.Token{
		{ Name: tokenizer.BraceOpen, Value: "{" },
		{ Name: tokenizer.DoubleQuotes, Value: "test" },
		{ Name: tokenizer.Colon, Value: ":" },
		{ Name: tokenizer.BraceOpen, Value: "{" },
		{ Name: tokenizer.DoubleQuotes, Value: "nested" },
		{ Name: tokenizer.Colon, Value: ":" },
		{ Name: tokenizer.DoubleQuotes, Value: "value" },
		{ Name: tokenizer.BraceClose, Value: "}" },
		{ Name: tokenizer.BraceClose, Value: "}" },
	}
	assertTokenSlicesAreEqual(t, got, want)
}

func TestJsonStringSuccessfullyTokenizedWithNestedObjectWithNumber(t *testing.T) {
	jsonString := `{"test": {"nested": 123}}`
	got := tokenizer.Tokenize(jsonString)
	want := []tokenizer.Token{
		{ Name: tokenizer.BraceOpen, Value: "{" },
		{ Name: tokenizer.DoubleQuotes, Value: "test" },
		{ Name: tokenizer.Colon, Value: ":" },
		{ Name: tokenizer.BraceOpen, Value: "{" },
		{ Name: tokenizer.DoubleQuotes, Value: "nested" },
		{ Name: tokenizer.Colon, Value: ":" },
		{ Name: tokenizer.Number, Value: "123" },
		{ Name: tokenizer.BraceClose, Value: "}" },
		{ Name: tokenizer.BraceClose, Value: "}" },
	}
	assertTokenSlicesAreEqual(t, got, want)
}

func TestJsonStringSuccessfullyTokenizedWithNestedObjectWithArry(t *testing.T) {
jsonString := `{"test": {"nested": [1,2,3]}}`
	got := tokenizer.Tokenize(jsonString)
	want := []tokenizer.Token{
		{ Name: tokenizer.BraceOpen, Value: "{" },
		{ Name: tokenizer.DoubleQuotes, Value: "test" },
		{ Name: tokenizer.Colon, Value: ":" },
		{ Name: tokenizer.BraceOpen, Value: "{" },
		{ Name: tokenizer.DoubleQuotes, Value: "nested" },
		{ Name: tokenizer.Colon, Value: ":" },
		{ Name: tokenizer.BracketOpen, Value: "[" },
		{ Name: tokenizer.Number, Value: "1" },
		{ Name: tokenizer.Comma, Value: "," },
		{ Name: tokenizer.Number, Value: "2" },
		{ Name: tokenizer.Comma, Value: "," },
		{ Name: tokenizer.Number, Value: "3" },
		{ Name: tokenizer.BracketClose, Value: "]" },
		{ Name: tokenizer.BraceClose, Value: "}" },
		{ Name: tokenizer.BraceClose, Value: "}" },
	}

	assertTokenSlicesAreEqual(t, got, want)
}


func assertTokenSlicesAreEqual(t *testing.T, got []tokenizer.Token, want []tokenizer.Token) {
	if len(got) != len(want) {
		t.Fatalf("got: %+q, want: %+q", got, want)
	}

	if(stringifyTokenSlice(got) != stringifyTokenSlice(want)) {
		t.Fatalf("got: %+q, want: %+q", got, want)
	}
}


func stringifyTokenSlice(tokens []tokenizer.Token) string {
	tokenString := ""
	for _, token := range tokens {
		tokenString += fmt.Sprintf("%+v\n", token)
	}
	return tokenString
}
