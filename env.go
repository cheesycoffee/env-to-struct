package envtostruct

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
)

// Set os env values to struct with "env:" tag. for slices use comma separated.
func Set(structField interface{}) (err error) {
	iType := reflect.TypeOf(structField)
	iValue := reflect.ValueOf(structField)

	if iValue.Kind() != reflect.Pointer {
		return errors.New("unable to set to non pointer struct")
	}

	iValue = iValue.Elem()
	iType = iType.Elem()

	if iType.Kind() != reflect.Struct {
		return errors.New("unable to parse non struct")
	}

	for i := 0; i < iType.NumField(); i++ {
		fieldType := iType.Field(i).Type
		fieldKind := fieldType.Kind()
		field := iValue.Field(i)

		// solve pointers
		if fieldKind == reflect.Pointer {
			fieldKind = fieldType.Elem().Kind()
			field.Set(reflect.New(field.Type().Elem()))
			field = field.Elem()
		}

		// handle embeded struct
		if fieldKind == reflect.Struct {
			if err := Set(field.Addr().Interface()); err != nil {
				return err
			}
			continue
		}

		envKey := iType.Field(i).Tag.Get("env")
		if envKey == "" || envKey == "-" {
			continue
		}
		envStringValue, ok := os.LookupEnv(envKey)
		if !ok {
			return fmt.Errorf(`undefined environment "%s" in .env file`, envKey)
		}

		if err := valueParser(fieldKind, field, envStringValue); err != nil {
			return err
		}

	}
	return nil
}

func valueParser(fieldKind reflect.Kind, field reflect.Value, stringValue string) (err error) {
	var fieldValue interface{}
	switch fieldKind {

	case reflect.String:
		fieldValue = stringValue

	case reflect.Bool:
		fieldValue, err = strconv.ParseBool(stringValue)

	case reflect.Int:
		fieldValue, err = strconv.Atoi(stringValue)

	case reflect.Int32:
		var i64 int64
		i64, err = strconv.ParseInt(stringValue, 10, 32)
		fieldValue = int32(i64)

	case reflect.Uint32:
		var ui64 uint64
		ui64, err = strconv.ParseUint(stringValue, 10, 32)
		fieldValue = uint32(ui64)

	case reflect.Int64:
		fieldValue, err = strconv.ParseInt(stringValue, 10, 64)

	case reflect.Uint64:
		fieldValue, err = strconv.ParseUint(stringValue, 10, 64)

	case reflect.Float32:
		var f64 float64
		f64, err = strconv.ParseFloat(stringValue, 32)
		fieldValue = float32(f64)

	case reflect.Float64:
		fieldValue, err = strconv.ParseFloat(stringValue, 64)

	case reflect.Slice:
		fieldValue, err = sliceParser(field.Type().String(), stringValue)

	case reflect.Map:
		fieldValue = reflect.MakeMap(field.Type())
		err = json.Unmarshal([]byte(stringValue), &fieldValue)

	default:
		err = fmt.Errorf("unsupported type %s", fieldKind.String())
	}

	if err != nil {
		return err
	}

	field.Set(reflect.ValueOf(fieldValue))

	return nil
}

func sliceParser(stringKind string, stringValue string) (val interface{}, err error) {
	fmt.Println(stringKind)
	arrString := strings.Split(strings.ReplaceAll(stringValue, " ", ""), ",")
	switch stringKind {
	case "[]string":
		return arrString, nil

	case "[]int":
		arrInt := []int{}
		for i := 0; i < len(arrString); i++ {
			fValue, err := strconv.Atoi(arrString[i])
			if err != nil {
				return arrInt, err
			}
			arrInt = append(arrInt, fValue)
		}
		return arrInt, nil

	case "[]int32":
		arrInt32 := []int32{}
		for i := 0; i < len(arrString); i++ {
			fValue, err := strconv.ParseInt(arrString[i], 10, 32)
			if err != nil {
				return arrInt32, err
			}
			arrInt32 = append(arrInt32, int32(fValue))
		}
		return arrInt32, nil

	case "[]uint32":
		arrUint32 := []uint32{}
		for i := 0; i < len(arrString); i++ {
			fValue, err := strconv.ParseUint(arrString[i], 10, 32)
			if err != nil {
				return arrUint32, err
			}
			arrUint32 = append(arrUint32, uint32(fValue))
		}
		return arrUint32, nil

	case "[]int64":
		arrInt64 := []int64{}
		for i := 0; i < len(arrString); i++ {
			fValue, err := strconv.ParseInt(arrString[i], 10, 64)
			if err != nil {
				return arrInt64, err
			}
			arrInt64 = append(arrInt64, fValue)
		}
		return arrInt64, nil

	case "[]uint64":
		arrUint64 := []uint64{}
		for i := 0; i < len(arrString); i++ {
			fValue, err := strconv.ParseUint(arrString[i], 10, 64)
			if err != nil {
				return arrUint64, err
			}
			arrUint64 = append(arrUint64, fValue)
		}
		return arrUint64, nil

	case "[]float32":
		arrFloat32 := []float32{}
		for i := 0; i < len(arrString); i++ {
			fValue, err := strconv.ParseFloat(arrString[i], 32)
			if err != nil {
				return arrFloat32, err
			}
			arrFloat32 = append(arrFloat32, float32(fValue))
		}
		return arrFloat32, nil

	case "[]float64":
		arrFloat64 := []float64{}
		for i := 0; i < len(arrString); i++ {
			fValue, err := strconv.ParseFloat(arrString[i], 64)
			if err != nil {
				return arrFloat64, err
			}
			arrFloat64 = append(arrFloat64, float64(fValue))
		}
		return arrFloat64, nil

	default:
		return nil, fmt.Errorf("unsupported slice type of %s", stringKind)
	}
}
