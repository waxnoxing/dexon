package errors

import (
	"fmt"
	"strconv"
	"strings"
)

// Error collects error information which should be reported to users.
type Error struct {
	// These keys are parts of SQL VM ABI. Database contract callers can
	// obtain values stored in these fields from function return values.
	Position uint32 // Position is the offset in bytes to the error location.
	Length   uint32 // Length is the length in bytes of the error token.
	Category ErrorCategory
	Code     ErrorCode

	// These keys are only used for debugging purposes and not included in ABI.
	// Values stored in these fields are not guaranteed to be stable, so they
	// MUST NOT be returned to the contract caller.
	Token   string // Token is the source code token where the error occurred.
	Prefix  string // Prefix identified the cause of the error.
	Message string // Message provides detailed the error message.
}

func (e Error) Error() string {
	b := strings.Builder{}
	// It is possible for an error to have zero length because not all errors
	// correspond to tokens. The parser can report an error with no length when
	// it encounters an unexpected token.
	if e.Position > 0 || e.Length > 0 {
		b.WriteString(fmt.Sprintf("offset %d", e.Position))
		if e.Length > 0 {
			b.WriteString(fmt.Sprintf(", length %d", e.Length))
		}
	} else {
		b.WriteString("unknown location")
	}
	b.WriteString(fmt.Sprintf(", category %d (%s), code %d (%s)",
		e.Category, e.Category, e.Code, e.Code))
	if e.Token != "" {
		b.WriteString(", token ")
		b.WriteString(strconv.Quote(e.Token))
	}
	if e.Prefix != "" {
		b.WriteString(", hint ")
		b.WriteString(strconv.Quote(e.Prefix))
	}
	if e.Message != "" {
		b.WriteString(", message: ")
		b.WriteString(e.Message)
	}
	return b.String()
}

// ErrorList is a list of Error.
type ErrorList []Error

func (e ErrorList) Error() string {
	b := strings.Builder{}
	for i := range e {
		if i > 0 {
			b.WriteByte('\n')
		}
		b.WriteString(e[i].Error())
	}
	return b.String()
}

// ErrorCategory is used to distinguish errors come from different phases.
type ErrorCategory uint16

// Error category starts from 1. Zero value is invalid.
const (
	ErrorCategoryNil ErrorCategory = iota
	ErrorCategoryLimit
	ErrorCategoryGrammar
	ErrorCategorySemantic
	ErrorCategoryRuntime
)

var errorCategoryMap = [...]string{
	ErrorCategoryLimit:    "limit",
	ErrorCategoryGrammar:  "grammar",
	ErrorCategorySemantic: "semantic",
	ErrorCategoryRuntime:  "runtime",
}

func (c ErrorCategory) Error() string {
	return errorCategoryMap[c]
}

// ErrorCode describes the reason of the error.
type ErrorCode uint16

// Error code starts from 1. Zero value is invalid.
const (
	ErrorCodeNil ErrorCode = iota
	ErrorCodeDepthLimitReached
	ErrorCodeParser
	ErrorCodeInvalidIntegerSyntax
	ErrorCodeInvalidNumberSyntax
	ErrorCodeIntegerOutOfRange
	ErrorCodeNumberOutOfRange
	ErrorCodeFractionalPartTooLong
	ErrorCodeEscapeSequenceTooShort
	ErrorCodeInvalidUnicodeCodePoint
	ErrorCodeUnknownEscapeSequence
	ErrorCodeInvalidBytesSize
	ErrorCodeInvalidIntSize
	ErrorCodeInvalidUintSize
	ErrorCodeInvalidFixedSize
	ErrorCodeInvalidUfixedSize
	ErrorCodeInvalidFixedFractionalDigits
	ErrorCodeInvalidUfixedFractionalDigits
	// Runtime Error
	ErrorCodeInvalidDataType
	ErrorCodeOverflow
	ErrorCodeUnderflow
	ErrorCodeIndexOutOfRange
	ErrorCodeInvalidCastType
	ErrorCodeDividedByZero
)

var errorCodeMap = [...]string{
	ErrorCodeDepthLimitReached:             "depth limit reached",
	ErrorCodeParser:                        "parser error",
	ErrorCodeInvalidIntegerSyntax:          "invalid integer syntax",
	ErrorCodeInvalidNumberSyntax:           "invalid number syntax",
	ErrorCodeIntegerOutOfRange:             "integer out of range",
	ErrorCodeNumberOutOfRange:              "number out of range",
	ErrorCodeFractionalPartTooLong:         "fractional part too long",
	ErrorCodeEscapeSequenceTooShort:        "escape sequence too short",
	ErrorCodeInvalidUnicodeCodePoint:       "invalid unicode code point",
	ErrorCodeUnknownEscapeSequence:         "unknown escape sequence",
	ErrorCodeInvalidBytesSize:              "invalid bytes size",
	ErrorCodeInvalidIntSize:                "invalid int size",
	ErrorCodeInvalidUintSize:               "invalid uint size",
	ErrorCodeInvalidFixedSize:              "invalid fixed size",
	ErrorCodeInvalidUfixedSize:             "invalid ufixed size",
	ErrorCodeInvalidFixedFractionalDigits:  "invalid fixed fractional digits",
	ErrorCodeInvalidUfixedFractionalDigits: "invalid ufixed fractional digits",
	// Runtime Error
	ErrorCodeInvalidDataType: "invalid data type",
	ErrorCodeOverflow:        "overflow",
	ErrorCodeUnderflow:       "underflow",
	ErrorCodeIndexOutOfRange: "index out of range",
	ErrorCodeInvalidCastType: "invalid cast type",
	ErrorCodeDividedByZero:   "divide by zero",
}

func (c ErrorCode) Error() string {
	return errorCodeMap[c]
}