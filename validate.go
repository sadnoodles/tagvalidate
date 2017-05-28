package main

import (
	"fmt"
	"reflect"
	"regexp"
	"strconv"
	"strings"

	"github.com/astaxie/beego/logs"
)

func regx_check(v string, regx string) bool {
	if match, err := regexp.MatchString(regx, v); err == nil {
		return match
	} else {
		logs.Info("Regexp error:", regx)
		return true
	}
}

var string_allowed = map[string](func(string, string) bool){
	"eq":        func(v string, t string) bool { return v == t },                      //(string)strictly equal to a string
	"neq":       func(v string, t string) bool { return v != t },                      //(string)strictly not equal to a string
	"starts":    func(v string, t string) bool { return strings.HasPrefix(v, t) },     //(string)strictly starts with a string
	"ends":      func(v string, t string) bool { return strings.HasSuffix(v, t) },     //(string)strictly ends with a string
	"contains":  func(v string, t string) bool { return strings.Contains(v, t) },      //(string)strictly contains a string
	"ncontains": func(v string, t string) bool { return !strings.Contains(v, t) },     //(string)strictly not contains a string
	"upper":     func(v string, t string) bool { return strings.ToUpper(v) == v },     //(bool) must be upper case
	"lower":     func(v string, t string) bool { return strings.ToLower(v) == v },     //(bool) must be lower case
	"empty":     func(v string, t string) bool { return (t == "false") && (v != "") }, //(bool) if allow empty?
	"len":       func(v string, t string) bool { return fmt.Sprint(len(v)) == t },     //(int) strictly set length to some value
	"max_len": func(v string, t string) bool {
		if ml, err := strconv.Atoi(t); err != nil {
			return ml > len(v)
		} else {
			return true
		}
	}, //(int) strictly set max length
	"min_len": func(v string, t string) bool {
		if ml, err := strconv.Atoi(t); err != nil {
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
	"email":                {0, ``},
	"url":                  {0, ``},
	"color":                {0, ``},
	"path":                 {0, ``},
	"uuid":                 {0, ``},
	"domain":               {1, ``},
	"date(%Y%M%D)":         {1, ``}, //allow template %Y%M%D-%h%m%s%c
	"json[(child)]":        {1, ``},
	"map[(child):(child)]": {1, ``}, // allow recursion
	"list[(child)]":        {1, ``}, // allow recursion
}

type FieldCheck struct {
}

func (checker *FieldCheck) checkBoolField(val reflect.Value, field reflect.StructField) error {
	var allowed = []string{} // No need to validate
	for _, tagname := range allowed {
		if tagvalue, ok := field.Tag.Lookup(tagname); ok {
			println(tagname, tagvalue, ok)
		} else {
			println(tagname, tagvalue, ok)
		}

	}
	return nil
}

func (checker *FieldCheck) checkBasicField(val reflect.Value, field reflect.StructField) error {
	return nil
}

func (checker *FieldCheck) checkStringField(val reflect.Value, field reflect.StructField) error {

	println(type_map)
	for tagname, tagfunc := range string_allowed {
		if tagvalue, ok := field.Tag.Lookup(tagname); ok {
			checked := tagfunc(val.String(), tagvalue)
			println(ok, tagname, tagvalue, val.String(), checked)
		} else {
			// println(ok, tagname, tagvalue)
		}

	}
	return nil
}

func (checker *FieldCheck) checkIntField(val reflect.Value, field reflect.StructField) error {
	var allowed = []string{
		"eq",   //(int/[]int)must in those values
		"neq",  //(int/[]int)must in those values
		"zero", //(bool) if allow zero?
		"max",  //(int) strictly set max
		"min",  //(int) strictly set min
		"func", //(string) given check func name under this struct
	}
	for _, tagname := range allowed {
		if tagvalue, ok := field.Tag.Lookup(tagname); ok {
			println(tagname, tagvalue, ok)
		} else {
			println(tagname, tagvalue, ok)
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
		if tagvalue, ok := field.Tag.Lookup(tagname); ok {
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
		err = checker.checkStringField(val, field) //todo
	case reflect.Int:
		err = checker.checkIntField(val, field) //todo
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

func validateTags(instance interface{}) []error {
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
		// field := st.Field(i)
		fieldvalue := val.Field(i)
		if !fieldvalue.CanInterface() {
			continue
		}
		// val := fieldvalue.Interface()
		// tag := field.Tag.Get("validate")
		// if tag == "" {
		// 	continue
		// }
		// vts := strings.Split(tag, ",")

		// for _, vt := range vts {
		// 	name := field.Name
		// 	if nameTag != "" {
		// 		name = field.Tag.Get(nameTag)
		// 	}

		// 	if len(prefix) > 0 {
		// 		name = prefix + "." + name
		// 	}

		// 	if vt == "struct" {
		// 		errs2 := v.validateAndTagPrefix(val, nameTag, name)
		// 		if len(errs2) > 0 {
		// 			errs = append(errs, errs2...)
		// 		}
		// 		continue
		// 	}

		// 	vf := v[vt]
		// 	if vf == nil {
		// 		errs = append(errs, BadField{
		// 			Field: name,
		// 			Err:   fmt.Errorf("undefined validator: %q", vt),
		// 		})
		// 		continue
		// 	}
		// 	if err := vf(val); err != nil {
		// 		errs = append(errs, BadField{name, err})
		// 	}
		// }
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
