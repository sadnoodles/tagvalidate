package tagvalidate

import (
	"fmt"
	"log"
	"reflect"
	"regexp"
	"strconv"
	"strings"
	"time"
)

const (
	Email          string = "^(((([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+(\\.([a-zA-Z]|\\d|[!#\\$%&'\\*\\+\\-\\/=\\?\\^_`{\\|}~]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])+)*)|((\\x22)((((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(([\\x01-\\x08\\x0b\\x0c\\x0e-\\x1f\\x7f]|\\x21|[\\x23-\\x5b]|[\\x5d-\\x7e]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(\\([\\x01-\\x09\\x0b\\x0c\\x0d-\\x7f]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}]))))*(((\\x20|\\x09)*(\\x0d\\x0a))?(\\x20|\\x09)+)?(\\x22)))@((([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|\\d|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.)+(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])|(([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])([a-zA-Z]|\\d|-|\\.|_|~|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])*([a-zA-Z]|[\\x{00A0}-\\x{D7FF}\\x{F900}-\\x{FDCF}\\x{FDF0}-\\x{FFEF}])))\\.?$"
	CreditCard     string = "^(?:4[0-9]{12}(?:[0-9]{3})?|5[1-5][0-9]{14}|6(?:011|5[0-9][0-9])[0-9]{12}|3[47][0-9]{13}|3(?:0[0-5]|[68][0-9])[0-9]{11}|(?:2131|1800|35\\d{3})\\d{11})$"
	ISBN10         string = "^(?:[0-9]{9}X|[0-9]{10})$"
	ISBN13         string = "^(?:[0-9]{13})$"
	UUID3          string = "^[0-9a-f]{8}-[0-9a-f]{4}-3[0-9a-f]{3}-[0-9a-f]{4}-[0-9a-f]{12}$"
	UUID4          string = "^[0-9a-f]{8}-[0-9a-f]{4}-4[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID5          string = "^[0-9a-f]{8}-[0-9a-f]{4}-5[0-9a-f]{3}-[89ab][0-9a-f]{3}-[0-9a-f]{12}$"
	UUID           string = "^[0-9a-f]{8}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{4}-[0-9a-f]{12}$"
	Alpha          string = "^[a-zA-Z]+$"
	Alphanumeric   string = "^[a-zA-Z0-9]+$"
	Numeric        string = "^[0-9]+$"
	Int            string = "^(?:[-+]?(?:0|[1-9][0-9]*))$"
	Float          string = "^(?:[-+]?(?:[0-9]+))?(?:\\.[0-9]*)?(?:[eE][\\+\\-]?(?:[0-9]+))?$"
	Hexadecimal    string = "^[0-9a-fA-F]+$"
	Hexcolor       string = "^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$"
	RGBcolor       string = "^rgb\\(\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*,\\s*(0|[1-9]\\d?|1\\d\\d?|2[0-4]\\d|25[0-5])\\s*\\)$"
	ASCII          string = "^[\x00-\x7F]+$"
	Multibyte      string = "[^\x00-\x7F]"
	FullWidth      string = "[^\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	HalfWidth      string = "[\u0020-\u007E\uFF61-\uFF9F\uFFA0-\uFFDC\uFFE8-\uFFEE0-9a-zA-Z]"
	Base64         string = "^(?:[A-Za-z0-9+\\/]{4})*(?:[A-Za-z0-9+\\/]{2}==|[A-Za-z0-9+\\/]{3}=|[A-Za-z0-9+\\/]{4})$"
	PrintableASCII string = "^[\x20-\x7E]+$"
	DataURI        string = "^data:.+\\/(.+);base64$"
	Latitude       string = "^[-+]?([1-8]?\\d(\\.\\d+)?|90(\\.0+)?)$"
	Longitude      string = "^[-+]?(180(\\.0+)?|((1[0-7]\\d)|([1-9]?\\d))(\\.\\d+)?)$"
	DNSName        string = `^([a-zA-Z0-9]{1}[a-zA-Z0-9_-]{0,62}){1}(\.[a-zA-Z0-9]{1}[a-zA-Z0-9_-]{1,62})*$`
	IP             string = `(([0-9a-fA-F]{1,4}:){7,7}[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,7}:|([0-9a-fA-F]{1,4}:){1,6}:[0-9a-fA-F]{1,4}|([0-9a-fA-F]{1,4}:){1,5}(:[0-9a-fA-F]{1,4}){1,2}|([0-9a-fA-F]{1,4}:){1,4}(:[0-9a-fA-F]{1,4}){1,3}|([0-9a-fA-F]{1,4}:){1,3}(:[0-9a-fA-F]{1,4}){1,4}|([0-9a-fA-F]{1,4}:){1,2}(:[0-9a-fA-F]{1,4}){1,5}|[0-9a-fA-F]{1,4}:((:[0-9a-fA-F]{1,4}){1,6})|:((:[0-9a-fA-F]{1,4}){1,7}|:)|fe80:(:[0-9a-fA-F]{0,4}){0,4}%[0-9a-zA-Z]{1,}|::(ffff(:0{1,4}){0,1}:){0,1}((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])|([0-9a-fA-F]{1,4}:){1,4}:((25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9])\.){3,3}(25[0-5]|(2[0-4]|1{0,1}[0-9]){0,1}[0-9]))`
	URLSchema      string = `((ftp|tcp|udp|wss?|https?):\/\/)`
	URLUsername    string = `(\S+(:\S*)?@)`
	Hostname       string = ``
	URLPath        string = `((\/|\?|#)[^\s]*)`
	URLPort        string = `(:(\d{1,5}))`
	URLIP          string = `([1-9]\d?|1\d\d|2[01]\d|22[0-3])(\.(1?\d{1,2}|2[0-4]\d|25[0-5])){2}(?:\.([0-9]\d?|1\d\d|2[0-4]\d|25[0-4]))`
	URLSubdomain   string = `((www\.)|([a-zA-Z0-9]([-\.][-\._a-zA-Z0-9]+)*))`
	URL            string = `^` + URLSchema + `?` + URLUsername + `?` + `((` + URLIP + `|(\[` + IP + `\])|(([a-zA-Z0-9]([a-zA-Z0-9-]+)?[a-zA-Z0-9]([-\.][a-zA-Z0-9]+)*)|(` + URLSubdomain + `?))?(([a-zA-Z\x{00a1}-\x{ffff}0-9]+-?-?)*[a-zA-Z\x{00a1}-\x{ffff}0-9]+)(?:\.([a-zA-Z\x{00a1}-\x{ffff}]{1,}))?))\.?` + URLPort + `?` + URLPath + `?$`
	SSN            string = `^\d{3}[- ]?\d{2}[- ]?\d{4}$`
	WinPath        string = `^[a-zA-Z]:\\(?:[^\\/:*?"<>|\r\n]+\\)*[^\\/:*?"<>|\r\n]*$`
	UnixPath       string = `^(/[^/\x00]*)+/?$`
	Semver         string = "^v?(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)\\.(?:0|[1-9]\\d*)(-(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*)(\\.(0|[1-9]\\d*|\\d*[a-zA-Z-][0-9a-zA-Z-]*))*)?(\\+[0-9a-zA-Z-]+(\\.[0-9a-zA-Z-]+)*)?$"
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
		var extra string
		ls := strings.SplitN(t, ",", 2)
		if len(ls) > 1 {
			t = ls[0]
			extra = ls[1]
		}
		if ta, ok := type_map[t]; ok {
			switch ta.CheckBy {
			case 0:
				return regx_check(v, ta.Regx)
			case 1:
				if extra == "" {
					extra = ta.Regx
				}
				if type_func, ok := type_func_map[t]; ok {
					return type_func(v, extra)
				}

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
	"int":      {0, `^(?:[-+]?(?:0|[1-9][0-9]*))$`},
	"float":    {0, `^(?:[-+]?(?:[0-9]+))?(?:\.[0-9]*)?(?:[eE][\+\-]?(?:[0-9]+))?$`},
	"hex":      {0, `^[a-fA-F0-9]+$`},
	"ipv4":     {0, `^((?:(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d)))\.){3}(?:25[0-5]|2[0-4]\d|((1\d{2})|([1-9]?\d))))$`}, //python reg
	"ip":       {0, IP},
	"email":    {0, Email},
	"url":      {0, `^https?\://[^\s\r\n]+$`},
	"hexcolor": {0, `^#?([0-9a-fA-F]{3}|[0-9a-fA-F]{6})$`},
	"fullpath": {0, "(" + WinPath + ")|(" + UnixPath + ")"},
	"uuid3":    {0, `^[a-f\d]{8}-[a-f\d]{4}-3[a-f\d]{3}-[a-f\d]{4}-[a-f\d]{12}$`},
	"uuid4":    {0, `^[a-f\d]{8}-[a-f\d]{4}-4[a-f\d]{3}-[89ab][a-f\d]{3}-[a-f\d]{12}$`},
	"uuid5":    {0, `^[a-f\d]{8}-[a-f\d]{4}-5[a-f\d]{3}-[89ab][a-f\d]{3}-[a-f\d]{12}$`},
	"uuid":     {0, `^[a-f\d]{8}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{4}-[a-f\d]{12}$`},
	"base64":   {0, Base64},                     //done
	"date":     {1, `2006-01-02T15:04:05.000Z`}, //done,allow template
	"json":     {1, ``},
	"domain":   {1, ``},
	"map":      {1, ``}, // allow recursion
	"list":     {1, ``}, // allow recursion
}

var type_func_map = map[string](func(string, string) bool){
	"domain": func(v string, t string) bool {
		return true
	},
	"date": func(v string, t string) bool {
		_, ok := time.Parse(t, v)
		return ok == nil

	},
}

func (checker *FieldCheck) checkStringField(val reflect.Value, field reflect.StructField) error {

	empty_tag := "empty"
	empty_value := ""
	valreal := val.String()
	table := string_allowed
	empty, _ := checker.getFieldTag(field, empty_tag)
	if empty == "true" && valreal == empty_value {
		return nil
	}
	for tagname, tagfunc := range table {
		if tagvalue, ok := field.Tag.Lookup(checker.getTagName(tagname)); ok {
			checked := tagfunc(valreal, tagvalue)
			// println(ok, tagname, tagvalue, val.String(), checked)
			if !checked {
				return GetError(3, field.Name, tagname, tagvalue, valreal)
			}
		}
	}
	return nil
}
