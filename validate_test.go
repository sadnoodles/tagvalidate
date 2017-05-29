package tagvalidate

import (
	"reflect"
	"testing"
)

type ValidateTestModel struct {
	Id          int64          `zero:"false"`
	Name        string         `empty:"false" type:"ip"`
	CustomCheck string         `func:"CheckCustom"`
	Actions     map[string]int `func:"CheckAction"`
	Emptytag    map[string]int
}

func (vtm *ValidateTestModel) CheckCustom(s string) bool {
	println("in custom func", s)
	return s == "hello"
}

func (vtm *ValidateTestModel) CheckAction(a map[string]int) bool {
	println("in custom func", a)
	return len(a) == 0
}

func TestFieldCheck_checkStringField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	type StringModel struct {
		Name string `type:"ip" starts:"xxx" regx:"hello?"`
	}
	var checker = new(FieldCheck)
	s := StringModel{Name: "293.2.3.4"}
	val := reflect.ValueOf(s)
	st := val.Type()
	i := 0
	fieldvalue := val.Field(i)

	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"checker_string_field", checker, args{fieldvalue, st.Field(i)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkStringField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkStringField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkByType(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}

	type StringModel struct {
		Id int32 `zero:"false" min:"1" max:"1"`
		// Id   int64  `zero:"false" min:"1" max:"1"`
		Name string `type:"ip" starts:"xxx" regx:"hello?"`
	}
	var checker = new(FieldCheck)
	s := StringModel{Name: "1.2.3.4"}
	val := reflect.ValueOf(s)
	st := val.Type()
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"checkByType", checker, args{val.Field(0), st.Field(0)}, false},
		{"checkByType", checker, args{val.Field(1), st.Field(1)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkByType(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkByType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_regx_check(t *testing.T) {
	type args struct {
		v    string
		regx string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
		{"reg_tests", args{"123123", `\d+`}, true},
		{"reg_tests", args{"asdas", `\d+`}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regx_check(tt.args.v, tt.args.regx); got != tt.want {
				t.Errorf("regx_check() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestFieldCheck_validateTags(t *testing.T) {
	type args struct {
		instance interface{}
	}
	var ins ValidateTestModel
	var chk = &FieldCheck{}
	ins.CustomCheck = "hello"
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
		{"full_field_check", chk, args{&ins},
			false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.validateTags(tt.args.instance); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.validateTags() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkBoolField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkBoolField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkBoolField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkBasicField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkBasicField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkBasicField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkIntField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkIntField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkIntField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkUintField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkUintField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkUintField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkFloatField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkFloatField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkFloatField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkStructField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkStructField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkStructField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkMapField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkMapField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkMapField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkPtrField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkPtrField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkPtrField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkSliceField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkSliceField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkSliceField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestFieldCheck_checkFuncField(t *testing.T) {
	type args struct {
		val   reflect.Value
		field reflect.StructField
	}
	tests := []struct {
		name    string
		checker *FieldCheck
		args    args
		wantErr bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkFuncField(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkFuncField() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestCheck(t *testing.T) {
	type args struct {
		instance interface{}
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
	}{
		// {"check instance",
		// 	args{
		// 		&ValidateTestModel{Id: 1, Name: "aaa", CustomCheck: "hello"}},
		// 	false},
		// {"check instance zero int",
		// 	args{
		// 		&ValidateTestModel{Name: "aaa", CustomCheck: "hello"}},
		// 	true},
		// {"check instance empty string",
		// 	args{
		// 		&ValidateTestModel{Id: 1, CustomCheck: "hello"}},
		// 	true},
		// {"check CustomCheck ref",
		// 	args{
		// 		&ValidateTestModel{Id: 1, Name: "aaa", CustomCheck: "he"}},
		// 	true},
		{"check instance ip",
			args{
				&ValidateTestModel{Id: 1, Name: "256.55.4.5", CustomCheck: "hello"}},
			false},
		{"check instance ip",
			args{
				&ValidateTestModel{Id: 1, Name: "255.22.33.11", CustomCheck: "hello"}},
			false},
		// {"check CustomCheck ptr",
		// 	args{
		// 		ValidateTestModel{Id: 1, Name: "aaa", CustomCheck: "he"}},
		// 	true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := Check(tt.args.instance); (err != nil) != tt.wantErr {
				t.Errorf("Check() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
