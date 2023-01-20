package envparse

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
)

var (
	defaultPrefix   = "APP"
	defaultMaxDepth = 100
)

// Parse scans through environment variables using mapping provided in interface
func Parse(ptr interface{}, envs []string) error {
	env := newEnvMap(defaultPrefix, envs)

	// Verify that ptr is a pointer to a struct
	interfaceValue := reflect.ValueOf(ptr)
	if ptr == nil || interfaceValue.Kind() != reflect.Ptr || interfaceValue.Elem().Kind() != reflect.Struct {
		return fmt.Errorf("interface must be a pointer to a struct")
	}

	structValue := interfaceValue.Elem()

	errorList := newErrorList()
	structValue.Set(parseStruct(structValue.Type(), env.GetPrefix(defaultPrefix), 0, errorList))
	if !errorList.IsEmpty() {
		return errorList
	}

	return nil
}

func parseStruct(structType reflect.Type, env envMap, depth int, errorList *ErrorList) reflect.Value {
	result := reflect.New(structType).Elem()

	depth++
	if depth >= defaultMaxDepth {
		errorList.Append(fmt.Errorf("too many levels of embedded structs, currently at depth '%d'", depth))
		return result
	}

	for i := 0; i < result.NumField(); i++ {
		fieldValue := result.Field(i)
		fieldType := structType.Field(i)
		if tagString, ok := fieldType.Tag.Lookup("env"); ok {
			if !fieldType.IsExported() {
				errorList.Append(fmt.Errorf("field '%s' is not exported, but containts 'env' tag", fieldType.Name))
				continue
			}

			tag, err := parseTag(tagString)
			if err != nil {
				errorList.Append(err)
				continue
			}

			fieldKind := fieldType.Type.Kind()
			if fieldKind == reflect.Ptr {
				if fieldValue.IsNil() {
					fieldValue.Set(reflect.New(fieldValue.Type().Elem()))
				}
				fieldValue = fieldValue.Elem()
			}

			if tag.required {
				if fieldValue.Type().Kind() == reflect.Struct {
					if !env.PrefixExists(tag.name) {
						// if field marked as required, but relevant environment variable is not provided,
						// that is parsing error and can skip further processing of this field.
						errorList.Append(fmt.Errorf("envvars under '%s' are required, but none provided", fmt.Sprintf("%s", tag.name)))
						continue
					}
				} else {
					if !env.Exists(tag.name) {
						// if field marked as required, but relevant environment variable is not provided,
						// that is parsing error and can skip further processing of this field.
						errorList.Append(fmt.Errorf("envvar '%s' is required, but not provided", fmt.Sprintf("%s", tag.name)))
						continue
					}
				}
			}

			fieldValueString := tag.defaultValue
			if env.Exists(tag.name) {
				fieldValueString = env.Get(tag.name)
			}

			err = switchFunc(fieldValue, env.GetPrefix(tag.name), fieldValueString, depth, errorList)
			if err != nil {
				errorList.Append(fmt.Errorf("error parsing field '%s': %w", fieldType.Name, err))
			}
		}
	}

	return result
}

func parseSlice(sliceType reflect.Type, env envMap, depth int, errorList *ErrorList) reflect.Value {
	result := reflect.MakeSlice(sliceType, 0, 0)

	depth++
	if depth >= defaultMaxDepth {
		errorList.Append(fmt.Errorf("too many levels of embedded structs, currently at depth '%d'", depth))
		return result
	}

	slicePrefixes := env.GetSlicePrefixes()
	if len(slicePrefixes) == 0 {
		return result
	}

	isPtr := false
	if sliceType.Elem().Kind() == reflect.Ptr {
		isPtr = true
	}

	makeType := sliceType.Elem()
	if isPtr {
		makeType = makeType.Elem()
	}

	var item, assignableItem reflect.Value

	for _, slicePrefix := range slicePrefixes {
		if isPtr {
			item = reflect.New(makeType)
			assignableItem = item.Elem()
		} else {
			item = reflect.New(makeType).Elem()
			assignableItem = item
		}

		err := switchFunc(assignableItem, env.GetPrefix(slicePrefix), slicePrefix, depth, errorList)
		if err != nil {
			errorList.Append(fmt.Errorf("error parsing slice: %w", err))
			break
		}

		result = reflect.Append(result, item)
	}

	return result
}

func parseInt(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(int(i))
}

func parseInt8(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseInt(val, 10, 8)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(int8(i))
}

func parseInt16(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseInt(val, 10, 16)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(int16(i))
}

func parseInt32(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseInt(val, 10, 32)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(int32(i))
}

func parseInt64(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseInt(val, 10, 64)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(i)
}

func parseUint(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(uint(i))
}

func parseUint8(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseUint(val, 10, 8)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(uint8(i))
}

func parseUint16(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseUint(val, 10, 16)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(uint16(i))
}

func parseUint32(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseUint(val, 10, 32)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(uint32(i))
}

func parseUint64(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseUint(val, 10, 64)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(i)
}

func parseFloat32(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseFloat(val, 32)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(float32(i))
}

func parseFloat64(val string, errorList *ErrorList) reflect.Value {
	if val == "" {
		val = "0"
	}
	i, err := strconv.ParseFloat(val, 64)
	if err != nil {
		errorList.Append(err)
	}
	return reflect.ValueOf(i)
}

func parseBool(val string, _ *ErrorList) reflect.Value {
	switch strings.ToLower(val) {
	case "false", "0", "f", "":
		return reflect.ValueOf(false)
	default:
		return reflect.ValueOf(true)
	}
}

func parseString(val string, _ *ErrorList) reflect.Value {
	return reflect.ValueOf(val)
}

func switchFunc(value reflect.Value, env envMap, valueString string, depth int, errorList *ErrorList) error {
	switch value.Kind() {
	case reflect.Int:
		value.Set(parseInt(valueString, errorList))
	case reflect.Int8:
		value.Set(parseInt8(valueString, errorList))
	case reflect.Int16:
		value.Set(parseInt16(valueString, errorList))
	case reflect.Int32:
		value.Set(parseInt32(valueString, errorList))
	case reflect.Int64:
		value.Set(parseInt64(valueString, errorList))
	case reflect.Uint:
		value.Set(parseUint(valueString, errorList))
	case reflect.Uint8:
		value.Set(parseUint8(valueString, errorList))
	case reflect.Uint16:
		value.Set(parseUint16(valueString, errorList))
	case reflect.Uint32:
		value.Set(parseUint32(valueString, errorList))
	case reflect.Uint64:
		value.Set(parseUint64(valueString, errorList))
	case reflect.Float32:
		value.Set(parseFloat32(valueString, errorList))
	case reflect.Float64:
		value.Set(parseFloat64(valueString, errorList))
	case reflect.Bool:
		value.Set(parseBool(valueString, errorList))
	case reflect.String:
		value.Set(parseString(valueString, errorList))
	case reflect.Slice:
		value.Set(parseSlice(value.Type(), env, depth+1, errorList))
	case reflect.Struct:
		value.Set(parseStruct(value.Type(), env, depth+1, errorList))
	default:
		return fmt.Errorf("unsupported type '%s'", value.Kind().String())
	}
	return nil
}
