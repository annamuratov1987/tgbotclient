package builder

import (
	"errors"
)

var (
	ErrorRequired                  = errors.New("Value is required")
	ErrorMismatchTypeInt64         = errors.New("Type must be of 'int64'")
	ErrorMismatchTypeString        = errors.New("Type must be of 'string'")
	ErrorMismatchTypeStringArray   = errors.New("Type must be of '[]string'")
	ErrorMismatchTypeStringOrInt64 = errors.New("Type must be of 'string' or 'int64'")
	ErrorMismatchTypeFileOrInt64   = errors.New("Type must be of 'os.File' or 'int64'")
)
