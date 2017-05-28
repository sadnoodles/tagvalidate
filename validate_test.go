package main

import (
	"reflect"
	"testing"
)

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
