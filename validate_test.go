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
	s := StringModel{Name: "1.2.3.4"}
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
		Name string `type:"ip" starts:"xxx" regx:"hello?"`
	}
	var checker = new(FieldCheck)
	s := StringModel{Name: "1.2.3.4"}
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
		{"checkByType", checker, args{fieldvalue, st.Field(i)}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := tt.checker.checkByType(tt.args.val, tt.args.field); (err != nil) != tt.wantErr {
				t.Errorf("FieldCheck.checkByType() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
