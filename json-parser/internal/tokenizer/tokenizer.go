package tokenizer

import (
	"fmt"
	"strings"
)


type Token struct {
  name TokenType
  value string
}

type TokenType string

const (
  BraceClose TokenType = "BraceClose"
  BraceOpen TokenType = "BraceOpen"
  BracketOpen TokenType = "BracketOpen"
  BracketClose TokenType = "BracketClose"
  Comma TokenType = "Comma"
  Colon TokenType = "Colon"
  Number TokenType = "Number"
  True TokenType = "True"
  False TokenType = "False"
  DoubleQuotes = "DoubleQuotes"
)

func Tokenize(jsonString string) []Token {
  current := 0
  tokens := []Token{}

  for index, char := range jsonString {
    switch char {
    case '{':
      tokens = append(tokens, Token{ name: BraceOpen, value: string(char)  }) 
    case '}':
      tokens = append(tokens, Token{ name: BraceClose, value: string(char)  }) 
    case '[':
      tokens = append(tokens, Token{ name: BracketOpen, value: string(char)  }) 
    case ']':
      tokens = append(tokens, Token{ name: BracketClose, value: string(char)  }) 
    case ',':
      tokens = append(tokens, Token{ name: Comma, value: string(char)  }) 
    case ':':
      tokens = append(tokens, Token{ name: Colon, value: string(char)  })
    case '"':
      tokens = append(tokens, Token{ name: DoubleQuotes , value: parseString(jsonString,index)  })
    }
    current++
  }
  fmt.Sprintf("tokens: ", tokens)

  return tokens
}

func parseString(jsonString string, startingIndex int) string {
  currentIndex := startingIndex
  slicedJson := jsonString[currentIndex:]
  fmt.Println("slicedJson " + slicedJson)
  stringValue := make([]string, 0)

  for _, char := range slicedJson {
    if char == '"' {
      break
    }
    stringValue = append(stringValue,string(char))
  }
  fmt.Println("stringValue", stringValue)
  return strings.Join(stringValue, "")
}
