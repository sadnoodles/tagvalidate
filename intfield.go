package validate

import (
	"log"
	"reflect"
	"strconv"
)

var int_allowed = map[string](func(int64, string) bool){
	"eq": func(v int64, t string) bool {
		if veq, err := strconv.Atoi(t); err == nil {
			return int64(veq) == v
		} else {
			return true
		}
	}, //(int/[]int)must in those values
	"neq": func(v int64, t string) bool {
		if veq, err := strconv.Atoi(t); err == nil {
			return int64(veq) != v
		} else {
			return true
		}
	}, //(int/[]int)must in those values
	"zero": func(v int64, t string) bool { return (t == "true") || ((t == "false") && (v != 0)) }, //(bool) if allow zero?
	"max": func(v int64, t string) bool {
		if ml, err := strconv.ParseInt(t, 10, 64); err == nil {
			return ml > v
		} else {
			return true
		}
	}, //(int) strictly set max
	"min": func(v int64, t string) bool {
		if ml, err := strconv.ParseInt(t, 10, 64); err == nil {
			return ml < v
		} else {
			log.Println("Convert error:", t)
			return true
		}
	}, //(int) strictly set min
	"func": func(v int64, t string) bool { return true }, //(string) given check func name under this struct
}

func (checker *FieldCheck) checkIntField(val reflect.Value, field reflect.StructField) error {

	empty_tag := "zero"
	empty_value := int64(0)
	valreal := val.Int()
	table := int_allowed
	empty, _ := checker.getFieldTag(field, empty_tag)
	if empty == "true" && valreal == empty_value {
		return nil
	}
	for tagname, tagfunc := range table {
		if tagvalue, ok := field.Tag.Lookup(checker.tag_prefix + tagname); ok {
			checked := tagfunc(valreal, tagvalue)
			// println(ok, tagname, tagvalue, val.String(), checked)
			if !checked {
				return GetError(3, field.Name, tagname, tagvalue, valreal)
			}
		}
	}
	return nil
}

func (checker *FieldCheck) checkUintField(val reflect.Value, field reflect.StructField) error {
	return nil
}
