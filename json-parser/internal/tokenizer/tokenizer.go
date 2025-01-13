package tokenizer

import (
	"strings"
	"regexp"
)


type Token struct {
  Name TokenType
  Value string
}

type TokenType string

var (
  BraceClose TokenType = "BraceClose"
  BraceOpen TokenType = "BraceOpen"
  BracketOpen TokenType = "BracketOpen"
  BracketClose TokenType = "BracketClose"
  Comma TokenType = "Comma"
  Colon TokenType = "Colon"
  Number TokenType = "Number"
  True TokenType = "True"
  False TokenType = "False"
  DoubleQuotes TokenType = "DoubleQuotes"
)

func Tokenize(jsonString string) []Token {
  tokens := []Token{}
	jsonLength := len(jsonString)
	regexToMatchNumber := regexp.MustCompile(`\d+`)

	for index := 0; index < jsonLength; index++ {
    switch jsonString[index] {
    case '{':
      tokens = append(tokens, Token{ Name: BraceOpen, Value: string(jsonString[index])  }) 
    case '}':
      tokens = append(tokens, Token{ Name: BraceClose, Value: string(jsonString[index])  }) 
    case '[':
      tokens = append(tokens, Token{ Name: BracketOpen, Value: string(jsonString[index])  }) 
    case ']':
      tokens = append(tokens, Token{ Name: BracketClose, Value: string(jsonString[index])  }) 
    case ',':
      tokens = append(tokens, Token{ Name: Comma, Value: string(jsonString[index])  }) 
    case ':':
      tokens = append(tokens, Token{ Name: Colon, Value: string(jsonString[index])  })
    case '"':
			stringValue := parseString(jsonString,index)
      tokens = append(tokens, Token{ Name: DoubleQuotes , Value: stringValue  })
		  index = index + len(stringValue) + 1
		case 't':
			if jsonString[index:index+4] == "true" {
				tokens = append(tokens, Token{ Name: True, Value: "true" })
				index = index + 3
			}
		case 'f':
			if jsonString[index:index+5] == "false" {
				tokens = append(tokens, Token{ Name: False, Value: "false" })
				index = index + 4
			}
    }

		if(regexToMatchNumber.MatchString(string(jsonString[index]))) {
			numberValue := parseNumber(jsonString,index)
			tokens = append(tokens, Token{ Name: Number, Value: numberValue  })
			index = index + len(numberValue) -1
		}

  }
  return tokens
}

func parseString(jsonString string, startingIndex int) string {
  currentIndex := startingIndex
	slicedJson := jsonString[currentIndex+1:]
  stringValue := make([]string, 0)

  for _, char := range slicedJson {
    if char == '"' {
      break
    }
    stringValue = append(stringValue,string(char))
  }
  return strings.Join(stringValue, "")
}


func parseNumber(jsonString string, startingIndex int) string {
	currentIndex := startingIndex
	slicedJson := jsonString[currentIndex:]
	numberValue := make([]string, 0)

	for _, char := range slicedJson {
		if char == ',' || char == '}' || char == ']' {
			break
		}
		numberValue = append(numberValue,string(char))
	}
	return strings.Join(numberValue, "")
}
