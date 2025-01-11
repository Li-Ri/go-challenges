package validator


type JsonValidator struct {
  jsonString string
}


func NewJsonValidator(jsonString string) JsonValidator {
  return JsonValidator{
    jsonString: jsonString,
    }
}


func (v JsonValidator) Validate() bool {
  if(v.jsonString == "{}") {
    return false
  }

  return true
}
