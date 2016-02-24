// Copyright (c) 2013 Kelsey Hightower. All rights reserved.
// Use of this source code is governed by the MIT License that can be found in
// the LICENSE file.

package envconfig

import (
	"errors"
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"syscall"
)

// ErrInvalidSpecification indicates that a specification is of the wrong type.
var ErrInvalidSpecification = errors.New("invalid specification must be a struct")

// A ParseError occurs when an environment variable cannot be converted to
// the type required by a struct field during assignment.
type ParseError struct {
	KeyName   string
	FieldName string
	TypeName  string
	Value     string
}

func (e *ParseError) Error() string {
	return fmt.Sprintf("envconfig.Process: assigning %[1]s to %[2]s: converting '%[3]s' to type %[4]s", e.KeyName, e.FieldName, e.Value, e.TypeName)
}

func Process(prefix string, spec interface{}) error {
	s := reflect.ValueOf(spec).Elem()
	if s.Kind() != reflect.Struct {
		return ErrInvalidSpecification
	}
	typeOfSpec := s.Type()
	for i := 0; i < s.NumField(); i++ {
		f := s.Field(i)
		if f.CanSet() {
			alt := typeOfSpec.Field(i).Tag.Get("envconfig")
			fieldName := typeOfSpec.Field(i).Name
			if alt != "" {
				fieldName = alt
			}
			key := strings.ToUpper(fmt.Sprintf("%s_%s", prefix, fieldName))
			// `os.Getenv` cannot differentiate between an explicitly set empty value
			// and an unset value. `os.LookupEnv` is preferred to `syscall.Getenv`,
			// but it is only available in go1.5 or newer.
			value, ok := syscall.Getenv(key)
			if !ok && alt != "" {
				key := strings.ToUpper(fieldName)
				value, ok = syscall.Getenv(key)
			}

			def := typeOfSpec.Field(i).Tag.Get("default")
			if def != "" && !ok {
				value = def
			}

			req := typeOfSpec.Field(i).Tag.Get("required")
			if value == "" {
				if req == "true" {
					return fmt.Errorf("required key %s missing value", key)
				}
				continue
			}

			parser := getParser(f.Type(), f.Kind())
			if parser == nil {
				continue
			}

			parsedValue, err := parser.Parse(value)
			if err != nil {
				return &ParseError{
					KeyName:   key,
					FieldName: fieldName,
					TypeName:  f.Type().String(),
					Value:     value,
				}
			}
			parser.Set(&f, parsedValue)

		}
	}
	return nil
}

func MustProcess(prefix string, spec interface{}) {
	if err := Process(prefix, spec); err != nil {
		panic(err)
	}
}

// fieldParser represents a parsing function (conversion from string) and a
// setting function (to set the value using reflection)
type fieldParser struct {
	Parse func(v string) (interface{}, error)
	Set   func(f *reflect.Value, v interface{})
}

// getParser returns back a FieldParser instance for the given type
func getParser(t reflect.Type, k reflect.Kind) (parser *fieldParser) {
	switch k {
	case reflect.String:
		parser = &fieldParser{
			Parse: func(v string) (interface{}, error) {
				return v, nil
			},
			Set: func(f *reflect.Value, v interface{}) {
				f.SetString(v.(string))
			},
		}
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		parser = &fieldParser{
			Parse: func(v string) (interface{}, error) {
				intValue, err := strconv.ParseInt(v, 0, t.Bits())
				if err != nil {
					return nil, err
				}

				return intValue, nil
			},
			Set: func(f *reflect.Value, v interface{}) {
				f.SetInt(v.(int64))
			},
		}
	case reflect.Bool:
		parser = &fieldParser{
			Parse: func(v string) (interface{}, error) {
				boolValue, err := strconv.ParseBool(v)
				if err != nil {
					return nil, err
				}
				return boolValue, nil
			},
			Set: func(f *reflect.Value, v interface{}) {
				f.SetBool(v.(bool))
			},
		}

	case reflect.Float32, reflect.Float64:
		parser = &fieldParser{
			Parse: func(v string) (interface{}, error) {
				floatValue, err := strconv.ParseFloat(v, t.Bits())
				if err != nil {
					return nil, err
				}
				return floatValue, nil
			},
			Set: func(f *reflect.Value, v interface{}) {
				f.SetFloat(v.(float64))
			},
		}
	case reflect.Slice:
		parser = &fieldParser{
			Parse: func(v string) (interface{}, error) {
				elemType := t.Elem()
				parser := getParser(elemType, elemType.Kind())
				strValues := strings.Split(v, ",")

				slice := reflect.MakeSlice(reflect.SliceOf(elemType), 0, 0)
				for i := range strValues {
					value, err := parser.Parse(strValues[i])
					if err != nil {
						return nil, err
					}

					itmValue := reflect.ValueOf(value)
					itmValue = itmValue.Convert(elemType)

					slice = reflect.Append(slice, itmValue)
				}
				return slice.Interface(), nil
			},
			Set: func(f *reflect.Value, v interface{}) {
				f.Set(reflect.ValueOf(v))
			},
		}
	}
	return
}
