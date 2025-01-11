package validator

import (
	"json-parser/internal/validator"
	"strconv"
	"testing"
)


func TestValidateInvalidJsonFalse(t *testing.T) {
  got := validator.NewJsonValidator("{}").Validate() 
  want := false

  if(got != want) {
    t.Fatalf(`Validate failed: got: %s, want: %s`, strconv.FormatBool(got), strconv.FormatBool(want))
  }
}

func TestValidateStringValueJsonAsTrue(t *testing.T) {
  got := validator.NewJsonValidator(`{"test":"value"}`).Validate()
  want := true

  if(got != want) {
    t.Fatalf(`Validate failed: got: %s, want: %s`, strconv.FormatBool(got), strconv.FormatBool(want))
  }
}
