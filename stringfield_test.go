package tagvalidate

import "testing"

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
	}
	for name, ta := range type_map {
		if ta.Regx == "" {
			continue
		}
		for _, str := range []string{
			"https://www.bbc.com",
			"https://www.bbc.com/",
			"/www.bbc.com/aaa",
			"www.ddd.cn",
			"a@sdoam.com",
			"c4b1866a-a021-41e1-8deb-cc9b6e9c1f36",
			"c4b1866acc9b6e9c1f36",
			"c4b1866a",
			"11222",
			"112..22",
			"11.2.22",
			".2.22",
			".2",
			"1.2",
			"-12",
			"-1.2",
			"汉纸的",
			"YQ==",
		} {
			tests = append(tests, struct {
				name string
				args args
				want bool
			}{name + "/" + str, args{str, ta.Regx}, false})
		}
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := regx_check(tt.args.v, tt.args.regx); got != tt.want {
				t.Errorf("regx_check() = %v, want %v", got, tt.want)
			}
		})
	}
}
