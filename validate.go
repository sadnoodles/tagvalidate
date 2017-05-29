package main

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
)

func regx_check(v string, regx string) bool {
	if match, err := regexp.MatchString(regx, v); err == nil {
		return match
	} else {
		log.Println("Regexp error:", regx)
		return true
	}
}

var string_allowed = map[string](func(string, string) bool){
	"eq":        func(v string, t string) bool { return v == t },                                         //(string)strictly equal to a string
	"neq":       func(v string, t string) bool { return v != t },                                         //(string)strictly not equal to a string
	"starts":    func(v string, t string) bool { return strings.HasPrefix(v, t) },                        //(string)strictly starts with a string
	"ends":      func(v string, t string) bool { return strings.HasSuffix(v, t) },                        //(string)strictly ends with a string
	"contains":  func(v string, t string) bool { return strings.Contains(v, t) },                         //(string)strictly contains a string
	"ncontains": func(v string, t string) bool { return !strings.Contains(v, t) },                        //(string)strictly not contains a string
	"upper":     func(v string, t string) bool { return strings.ToUpper(v) == v },                        //(bool) must be upper case
	"lower":     func(v string, t string) bool { return strings.ToLower(v) == v },                        //(bool) must be lower case
	"empty":     func(v string, t string) bool { return (t == "true") || ((t == "false") && (v != "")) }, //(bool) if allow empty?
	"len":       func(v string, t string) bool { return fmt.Sprint(len(v)) == t },                        //(int) strictly set length to some value
	"max_len": func(v string, t string) bool {
		if ml, err := strconv.Atoi(t); err == nil {
			return ml > len(v)
		} else {
			return true
		}
	}, //(int) strictly set max length
	"min_len": func(v string, t string) bool {
		if ml, err := strconv.Atoi(t); err == nil {
			return ml < len(v)
		} else {
			return true
		}
	}, //(int) strictly set min length
	"type": func(v string, t string) bool {
		if ta, ok := type_map[t]; ok {
			if ta.CheckBy == 0 {
				return regx_check(v, ta.Regx)
			}
		}
		return true
	}, //(string)validate as specific type. e.g:int, float, hex, bin, ip, email, url, uuid, date, month, domain
	"regx": regx_check,                                    //(string) only allowed matched strings
	"func": func(v string, t string) bool { return true }, //(string) given check func name under this struct
}

type typeAttr struct {
	CheckBy int    //0:regx, 1:func
	Regx    string //
}

var type_map = map[string]typeAttr{
	"int":                  {0, `^\d+$`},
	"float":                {0, `^(\d+)?\.\d+$`},
	"hex":                  {0, `^[a-fA-F0-9]+$`},
	"bin":                  {0, `^[01]+$`},
	"ip":                   {0, `^((?:(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d))))$`}, //python reg
	"base64":               {0, ``},
	"email":                {0, ``},
	"url":                  {0, ``},
	"color":                {0, ``},
	"path":                 {0, ``},
	"uuid":                 {0, `^[a-f\d]{8}-([a-f\d]{4}-){3}[a-f\d]{12}$`},
	"domain":               {1, ``},
	"date(%Y%M%D)":         {1, ``}, //allow template %Y%M%D-%h%m%s%c
	"json[(child)]":        {2, ``},
	"map[(child):(child)]": {2, ``}, // allow recursion
	"list[(child)]":        {2, ``}, // allow recursion
}

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

type FieldCheck struct {
	tag_prefix string
}

func (checker *FieldCheck) checkBoolField(val reflect.Value, field reflect.StructField) error {
	return nil
}

func (checker *FieldCheck) checkBasicField(val reflect.Value, field reflect.StructField) error {
	return nil
}

func reflectCall(ins interface{}, funcname string, arg1 reflect.Value) bool {
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
	if method := reflect.ValueOf(ins).MethodByName(funcname); method.IsValid() {
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

func (checker *FieldCheck) checkStringField(val reflect.Value, field reflect.StructField) error {

	for tagname, tagfunc := range string_allowed {
		if tagvalue, ok := field.Tag.Lookup(checker.tag_prefix + tagname); ok {
			checked := tagfunc(val.String(), tagvalue)
			println(ok, tagname, tagvalue, val.String(), checked)
		} else {
			// println(ok, tagname, tagvalue)
		}

	}
	return nil
}

func (checker *FieldCheck) checkIntField(val reflect.Value, field reflect.StructField) error {

	for tagname, tagfunc := range int_allowed {
		if tagvalue, ok := field.Tag.Lookup(checker.tag_prefix + tagname); ok {
			checked := tagfunc(val.Int(), tagvalue)
			println(ok, tagname, tagvalue, val.String(), checked)
		} else {
			// println(ok, tagname, tagvalue)
		}

	}
	return nil
}

func (checker *FieldCheck) checkUintField(val reflect.Value, field reflect.StructField) error {
	return nil
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

func (checker *FieldCheck) validateTags(instance interface{}) []error {
	val := reflect.ValueOf(instance)
	if val.Kind() == reflect.Ptr {
		val = val.Elem()
	}

	st := val.Type()
	if st == nil || st.Kind() != reflect.Struct {
		return nil
	}

	var errs []error

	for i := 0; i < st.NumField(); i++ {
		field := st.Field(i)
		fieldvalue := val.Field(i)
		if !fieldvalue.CanInterface() {
			continue
		}
		if tagvalue, ok := field.Tag.Lookup(checker.tag_prefix + "func"); ok {
			var checked bool
			checked = reflectCall(instance, tagvalue, fieldvalue)
			if !checked {
				err1 := fmt.Errorf("Error checking: field: %s, value: %#v, validate use: func", field.Name, fieldvalue.Interface())
				errs = append(errs, err1)
			}
		}
		err := checker.checkByType(fieldvalue, field)
		if err != nil {
			errs = append(errs, err)
		}
	}

	return errs
}

func main() {
	type S struct {
		F    string  `species:"gopher" color:"blue"`
		Id   int64   `min:"1" max:"5" eq:"[4]" neq:"[4,5]" zero:"false" func:"check_id"`
		Time float64 `min:"1" max:"5" eq:"[4]" neq:"[4,5]" zero:"false" func:"check_time"`
		Date string  `min:"1" max:"5" eq:"[4]" neq:"[4,5]" zero:"false" func:"check_date"`
		Name string  `eq:"ab" neq:"" starts:"x" type:"" ends:"g" contains:"s" upper:"false" lower:"false" regx:"" empty:"false" len:"2" max_len:"3" min_len:"1" func:"check_name"`

		// type includes:int, float, hex, bin, ip, email, url, date, month, domain
	}

	s := S{}
	st := reflect.TypeOf(s)
	field := st.Field(0)
	field1 := st.Field(4)

	fmt.Println(field.Tag.Get("color"), field.Tag.Get("species"))
	fmt.Println(field1.Type)
	fmt.Println(field1.Tag.Get("starts"))
	fmt.Println(field1.Tag.Get("lenss"))

	for i := 0; i < st.NumField(); i++ {
		f := st.Field(i)
		// fv := val.Field(i)
		// if !fv.CanInterface() {
		// 	continue
		// }
		// val := fv.Interface()
		tag := f.Tag.Get("validate")
		println(tag)

	}
}
