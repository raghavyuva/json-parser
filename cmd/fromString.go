/*
Copyright © 2025 raghavyuva <raghavyuva@gmail.com>
*/
package cmd

import (
	"errors"
	"fmt"

	"github.com/spf13/cobra"
)

// list of variables used to determine each characters / symbols for readablity
var (
	COMMA        = ','
	COLON        = ':'
	LEFTBRACKET  = '['
	RIGHTBRACKET = ']'
	LEFTBRACE    = '{'
	RIGHTBRACE   = '}'
	QUOTE        = '"'
)

var (
	errEmptyStringFound             = errors.New("empty string found")
	errInvalidOpeningQuotesInString = errors.New("invalid opening quotes in string")
	errInvalidEndingQuotesInString  = errors.New("invalid ending quotes in string")
	errFoundQuotesInBetween         = errors.New("found quotes in between string")
	errInvalidDigitFound            = errors.New("invalid digit found")
	errInvalidBooleanFound          = errors.New("invalid boolean found")
	errInvalidNullValue             = errors.New("invalid null value")
)

// fromStringCmd represents the fromString command
var fromStringCmd = &cobra.Command{
	Use:   "fromString [json_string]",
	Short: "converts string to json",
	Long: `fromString converts string to json and prints out the result.
 Example: json-parser fromString '{"name": "Tony Stark", "age": 22}'`,
	Args: cobra.ExactArgs(1),
	RunE: func(cmd *cobra.Command, args []string) error {
		tokens, err := lexer(args[0])
		if err != nil {
			return fmt.Errorf("failed to parse json: %v", err)
		}

		fmt.Printf("Tokens: %v\n", tokens)
		return nil
	},
}

func init() {
	rootCmd.AddCommand(fromStringCmd)
}

// lexer is a function that takes a string as input and returns a token or an error
// typical use of lexer is to convert a string to a list of tokens
// Example input : '{"foo": [1, 2, {"bar": 2}]}'
// Example output : ['{', 'foo', ':', '[', 1, ',', 2, ',', '{', 'bar', ':', 2, '}', ']', '}']
func lexer(s string) ([]string, error) {
	tokens := []string{}
	i := 0
	length := len(s)

	for i < length {
		char := s[i]

		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			i++
			continue
		}

		if char == '{' || char == '}' || char == '[' || char == ']' || char == ':' || char == ',' {
			tokens = append(tokens, string(char))
			i++
			continue
		}

		if char == '"' {
			start := i
			i++
			for i < length && s[i] != '"' {
				if s[i] == '\\' {
					i += 2
					continue
				}
				i++
			}
			if i >= length {
				return nil, errInvalidEndingQuotesInString
			}
			i++

			token, err := lex_string(s[start:i])
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			continue
		}

		if char == '-' || char == '+' || (char >= '0' && char <= '9') {
			start := i
			i++
			for i < length && (s[i] == '.' || s[i] == 'e' || s[i] == '+' || s[i] == '-' || (s[i] >= '0' && s[i] <= '9')) {
				i++
			}
			token, err := lex_number(s[start:i])
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			continue
		}

		if i+4 <= length && s[i:i+4] == "true" {
			token, err := lex_boolean("true")
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			i += 4
			continue
		}
		if i+5 <= length && s[i:i+5] == "false" {
			token, err := lex_boolean("false")
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			i += 5
			continue
		}

		if i+4 <= length && s[i:i+4] == "null" {
			token, err := lex_null_value("null")
			if err != nil {
				return nil, err
			}
			tokens = append(tokens, token)
			i += 4
			continue
		}

		return nil, fmt.Errorf("invalid character at position %d: %c", i, char)
	}

	return tokens, nil
}

// lex_string validates and extracts the content inside quotes from a given string.
// It returns a pointer to the extracted string content or an error if the input string is malformed.
// Errors returned can include:
// - errEmptyStringFound: if the input string is empty
// - errInvalidOpeningQuotesInString: if the input string does not start with opening quotes
// - errInvalidEndingQuotesInString: if the input string does not end with closing quotes
// - errFoundQuotesInBetween: if the input string contains quotes in between
func lex_string(s string) (string, error) {
	output := ""
	string_length := len(s)

	if s == "" {
		return "", errEmptyStringFound
	}

	if s[0] != byte(QUOTE) {
		return "", errInvalidOpeningQuotesInString
	}

	if s[string_length-1] != byte(QUOTE) {
		return "", errInvalidEndingQuotesInString
	}

	for index, character := range s {
		if index == 0 || index == string_length-1 {
			continue
		}

		if character == QUOTE {
			return "", errFoundQuotesInBetween
		}

		output += string(character)
	}

	return output, nil
}

// lex_number validates and extracts the number from a given string.
// It returns the extracted number as a string or an error if the input string is malformed.
// Errors returned can include:
// - errEmptyStringFound: if the input string is empty
// - errInvalidDigitFound: if the input string contains invalid digits
func lex_number(s string) (string, error) {
	output := ""
	if s == "" {
		return "", errEmptyStringFound
	}

	for _, character := range s {
		if (character < '0' || character > '9') && character != '+' && character != '.' && character != '-' && character != 'e' {
			return "", errInvalidDigitFound
		}
		output += string(character)
	}
	return output, nil
}

// lex_boolean validates and extracts the boolean from a given string.
// It returns the extracted boolean as a string or an error if the input string is malformed.
// Errors returned can include:
// - errInvalidBooleanFound: if the input string does not contain a valid boolean
func lex_boolean(s string) (string, error) {
	output := ""

	if s == "true" || s == "false" {
		output = s
		return output, nil
	}

	return "", errInvalidBooleanFound
}

// lex_null_value validates and extracts the null value from a given string.
// It returns the extracted null value as a string or an error if the input string is malformed.
// Errors returned can include:
// - errInvalidNullValue: if the input string does not contain a valid null value
func lex_null_value(s string) (string, error) {
	if s == "null" {
		return s, nil
	}

	return "", errInvalidNullValue
}
