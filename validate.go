package tagvalidate

import (
	"fmt"
	"log"
	"reflect"
)

type FieldCheck struct {
	tag_prefix string
}

var msgmap = map[int]string{
	0: "",
	1: "Must be a struct!",
	2: "Custom check func didn't pass.",
	3: "Validate error.",
}

func GetError(id int, fieldname string, tagname string, tagvalve string, gotvalue interface{}) error {
	var err error
	pos := fmt.Sprintf("Field:%s tag:%s value:%s, Got: %#v", fieldname, tagname, tagvalve, gotvalue)
	if id == 0 {
		return nil
	}
	err = fmt.Errorf("%s @%s", msgmap[id], pos)
	return err
}

func Check(instance interface{}) error {
	var err error
	ck := new(FieldCheck)
	err = ck.ValidateStruct(instance)
	return err
}

func (checker *FieldCheck) checkBoolField(val reflect.Value, field reflect.StructField) error {
	return nil
}

func (checker *FieldCheck) checkBasicField(val reflect.Value, field reflect.StructField) error {
	return nil
}

func reflectCall(ins interface{}, funcname string, arg1 reflect.Value) bool {
	return reflectCallV(reflect.ValueOf(ins), funcname, arg1)
}

func reflectCallV(ins reflect.Value, funcname string, arg1 reflect.Value) bool {
	var ret bool
	ret = true
	defer func() {
		e := recover()
		if e != nil {
			if ee, ok := e.(error); ok {
				log.Println(ee.Error())
			}
		}
	}()
	if method := ins.MethodByName(funcname); method.IsValid() {
		mtype := method.Type()
		args := make([]reflect.Value, mtype.NumIn())
		args[0] = arg1
		results := method.Call(args)
		if len(results) > 0 {
			ret = results[0].Bool()
		}

	}
	return ret

}

func (checker *FieldCheck) getFieldTag(field reflect.StructField, tagname string) (string, bool) {
	if tagvalue, ok := field.Tag.Lookup(checker.tag_prefix + tagname); ok {
		return tagvalue, true
	}
	return "", false
}

func (checker *FieldCheck) checkFloatField(val reflect.Value, field reflect.StructField) error {
	var allowed = []string{
		"eq",   //(float/[]float)must in those values
		"neq",  //(float/[]float)must in those values
		"zero", //(bool) if allow zero?
		"max",  //(float) strictly set max
		"min",  //(float) strictly set min
		"func", //(string) given check func name under this struct
	}
	for _, tagname := range allowed {
		if tagvalue, ok := field.Tag.Lookup(checker.tag_prefix + tagname); ok {
			println(tagname, tagvalue, ok)
		} else {
			println(tagname, tagvalue, ok)
		}

	}
	return nil
}

func (checker *FieldCheck) checkStructField(val reflect.Value, field reflect.StructField) error {
	return nil
}

func (checker *FieldCheck) checkMapField(val reflect.Value, field reflect.StructField) error {
	return nil
}

func (checker *FieldCheck) checkPtrField(val reflect.Value, field reflect.StructField) error {
	return nil
}

func (checker *FieldCheck) checkSliceField(val reflect.Value, field reflect.StructField) error {
	return nil
}

func (checker *FieldCheck) checkFuncField(val reflect.Value, field reflect.StructField) error {
	return nil
}

func getKind(val reflect.Value) reflect.Kind {
	kind := val.Kind()

	switch {
	case kind >= reflect.Int && kind <= reflect.Int64:
		return reflect.Int
	case kind >= reflect.Uint && kind <= reflect.Uint64:
		return reflect.Uint
	case kind >= reflect.Float32 && kind <= reflect.Float64:
		return reflect.Float32
	default:
		return kind
	}
}

func (checker *FieldCheck) checkByType(val reflect.Value, field reflect.StructField) error {
	var err error

	dataKind := getKind(val)
	switch dataKind {
	case reflect.Bool:
		err = checker.checkBoolField(val, field) //todo
	case reflect.Interface:
		err = checker.checkBasicField(val, field)
	case reflect.String:
		err = checker.checkStringField(val, field) // done
	case reflect.Int:
		err = checker.checkIntField(val, field) // done
	case reflect.Uint:
		err = checker.checkUintField(val, field) //todo
	case reflect.Float32:
		err = checker.checkFloatField(val, field) //todo
	case reflect.Struct:
		err = checker.checkStructField(val, field)
	case reflect.Map:
		err = checker.checkMapField(val, field)
	case reflect.Ptr:
		err = checker.checkPtrField(val, field)
	case reflect.Slice:
		err = checker.checkSliceField(val, field)
	case reflect.Func:
		err = checker.checkFuncField(val, field)
	default:
		return fmt.Errorf("Unsupported type: %s", dataKind)
	}

	return err
}

func (checker *FieldCheck) ValidateStructV(val reflect.Value) error {
	// https://stackoverflow.com/questions/24348184/get-pointer-to-value-using-reflection
	if val.Kind() == reflect.Interface && !val.IsNil() {
		elm := val.Elem()
		if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
			val = elm
		}
	}
	oldval := val
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	for i := 0; i < val.NumField(); i++ {
		valueField := val.Field(i)
		typeField := val.Type().Field(i)
		address := "not-addressable"

		if valueField.Kind() == reflect.Interface && !valueField.IsNil() {
			elm := valueField.Elem()
			if elm.Kind() == reflect.Ptr && !elm.IsNil() && elm.Elem().Kind() == reflect.Ptr {
				valueField = elm
			}
		}

		if valueField.Kind() == reflect.Ptr {
			valueField = valueField.Elem()

		}
		if valueField.CanAddr() {
			address = fmt.Sprintf("0x%X", valueField.Addr().Pointer())
		}

		fmt.Printf("Field Name: %s,\t Field Value: %v,\t Address: %v\t, Field type: %v\t, Field kind: %v\n", typeField.Name,
			valueField.Interface(), address, typeField.Type, valueField.Kind())

		if valueField.Kind() == reflect.Struct {
			return checker.ValidateStructV(valueField)
		}
		if tagvalue, ok := typeField.Tag.Lookup(checker.tag_prefix + "func"); ok {
			var checked bool
			checked = reflectCallV(oldval, tagvalue, valueField)
			if !checked {
				return GetError(2, typeField.Name, "func", tagvalue, valueField.Interface())
			}
		}

		err := checker.checkByType(valueField, typeField)
		if err != nil {
			return err
		}
	}
	return nil
}

func (checker *FieldCheck) ValidateStruct(v interface{}) error {
	return checker.ValidateStructV(reflect.ValueOf(v))
}
