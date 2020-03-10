package vartable

import (
	"reflect"
	"testing"
)

func Test_splitFields(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		{
			"test1",
			args{
				"hello",
			},
			[]string{"hello"},
			false,
		},
		{
			"test2",
			args{
				"hello.tourist",
			},
			[]string{"hello", "tourist"},
			false,
		},
		{
			"test2",
			args{
				"hello[tourist]",
			},
			[]string{"hello", "tourist"},
			false,
		},
		{
			"test3",
			args{
				"hello.foo[tourist]",
			},
			[]string{"hello", "foo", "tourist"},
			false,
		},
		{
			"test4",
			args{
				"hello.foo.tourist",
			},
			[]string{"hello", "foo", "tourist"},
			false,
		},
		{
			"test5",
			args{
				"hello.foo[tourist.dubist]",
			},
			[]string{"hello", "foo", "tourist.dubist"},
			false,
		},
		{
			"test6",
			args{
				"hello.foo[tourist.dubist[in]]",
			},
			[]string{"hello", "foo", "tourist.dubist[in]"},
			false,
		},
		{
			"test7",
			args{
				"hello.foo[tourist.dubist[in.bp[capitol]]]",
			},
			[]string{"hello", "foo", "tourist.dubist[in.bp[capitol]]"},
			false,
		},
		{
			"test8",
			args{
				"hello.foo[tourist.dubist[in.bp[capitol]]].for.an.little",
			},
			[]string{"hello", "foo", "tourist.dubist[in.bp[capitol]]", "for", "an", "little"},
			false,
		},
		{
			"test9",
			args{
				"hello.foo[in][bp]",
			},
			[]string{"hello", "foo", "in", "bp"},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := splitFields(tt.args.path)
			if (err != nil) != tt.wantErr {
				t.Errorf("splitFields() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("splitFields() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_unflattenPath(t *testing.T) {
	type args struct {
		fields []interface{}
		value  interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			"test1",
			args{
				[]interface{}{"hello", "tourist"},
				"okay",
			},
			map[string]interface{}{
				"hello": map[interface{}]interface{}{
					"tourist": "okay",
				},
			},
			false,
		},
		{
			"test2",
			args{
				[]interface{}{"hello", "tourist", "dubist"},
				"okay",
			},
			map[string]interface{}{
				"hello": map[interface{}]interface{}{
					"tourist": map[interface{}]interface{}{
						"dubist": "okay",
					},
				},
			},
			false,
		},
		{
			"test2",
			args{
				[]interface{}{"hello", "tourist", "dubist.in.bp"},
				"okay",
			},
			map[string]interface{}{
				"hello": map[interface{}]interface{}{
					"tourist": map[interface{}]interface{}{
						"dubist.in.bp": "okay",
					},
				},
			},
			false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := unflattenPath(tt.args.fields, tt.args.value)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("unflattenPath() = %v, want %v", got, tt.want)
			}
		})
	}
}
