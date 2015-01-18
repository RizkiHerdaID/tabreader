// GPL v3.0

// Package tabreader implements a reader for fixed-width fields.
// It is based on the code from kopiczko in this StackOverflow question:
// http://stackoverflow.com/questions/27968385/reading-tabular-data-with-fixed-width-and-missing-values
package tabreader

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

type scanner struct {
	len   int
	parts []int
}

// New returns a new scanner that can be used to scan lines. parts is a variadic with the sequential lengths of each field you want to scan.
func New(parts ...int) *scanner {
	len := 0
	for _, v := range parts {
		len += v
	}
	return &scanner{len, parts}
}

// Scan parses a line according to the provided scanner specification. Arguments types are discovered using reflect. Valid types are: int, int32, int64, float32, float64.
// If a field is empty (all blanks), the zero-value for its type is used. line length must be at least the sum of all field lengths.
// Number of arguments must match the number of fields used when creating the scanner.
//
// Scan returns the number of items read. If lower than the number of fields, an error is also returned.
func (s *scanner) Scan(line string, args ...interface{}) (n int, err error) {
	if i := len(line); i < s.len {
		return 0, fmt.Errorf("exepected string of size at least %d, actual %d", s.len, i)
	}
	if len(args) != len(s.parts) {
		return 0, fmt.Errorf("expected %d args, actual %d", len(s.parts), len(args))
	}
	n = 0
	start := 0
	for ; n < len(args); n++ {
		a := args[n]
		l := s.parts[n]
		if err = scanOne(line[start:start+l], a); err != nil {
			return
		}
		start += l
	}
	return n, nil
}

func scanOne(s string, arg interface{}) (err error) {
	s = strings.TrimSpace(s)
	switch v := arg.(type) {
	case *int:
		if s == "" {
			*v = int(0)
		} else {
			*v, err = strconv.Atoi(s)
		}
	case *int32:
		if s == "" {
			*v = int32(0)
		} else {
			var val int64
			val, err = strconv.ParseInt(s, 10, 32)
			*v = int32(val)
		}
	case *int64:
		if s == "" {
			*v = int64(0)
		} else {
			*v, err = strconv.ParseInt(s, 10, 64)
		}
	case *float32:
		if s == "" {
			*v = float32(0)
		} else {
			var val float64
			val, err = strconv.ParseFloat(s, 32)
			*v = float32(val)
		}
	case *float64:
		if s == "" {
			*v = float64(0)
		} else {
			*v, err = strconv.ParseFloat(s, 64)
		}
	default:
		val := reflect.ValueOf(v)
		err = fmt.Errorf("can't scan type: " + val.Type().String())
	}
	return
}
