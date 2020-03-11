package maptable

import (
	"reflect"
	"testing"
)

func testAddPathHelper(original map[string]interface{}, path string, value interface{}) (map[string]interface{}, error) {
	vt := New()
	vt.vars = original
	err := vt.AddPath(path, value)
	if err != nil {
		return nil, err
	}
	return vt.vars, err
}

func Test_test_AddPath_Helper(t *testing.T) {
	type args struct {
		original map[string]interface{}
		path     string
		value    interface{}
	}
	tests := []struct {
		name    string
		args    args
		want    map[string]interface{}
		wantErr bool
	}{
		{
			name: "test1",
			args: args{
				original: map[string]interface{}{
					"date":  "wednesday",
					"hello": "foo",
				},
				path:  "hello",
				value: "bar",
			},
			want: map[string]interface{}{
				"hello": "bar",
				"date":  "wednesday",
			},
		},
		{
			name: "test2",
			args: args{
				original: map[string]interface{}{
					"foo": "bar",
					"hello": map[interface{}]interface{}{
						"user": map[interface{}]interface{}{
							"name": "john",
						},
					},
				},
				path:  "hello.user.name.yolo",
				value: "new",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "test3",
			args: args{
				original: map[string]interface{}{
					"hello": map[interface{}]interface{}{
						"foo": "bar",
						"user": map[interface{}]interface{}{
							"name": "john",
						},
					},
				},
				path:  "hello.user",
				value: "jane",
			},
			want: map[string]interface{}{
				"hello": map[interface{}]interface{}{
					"user": "jane",
					"foo":  "bar",
				},
			},
			wantErr: false,
		},
		{
			name: "test4",
			args: args{
				original: map[string]interface{}{
					"hello": map[interface{}]interface{}{
						"user": map[interface{}]interface{}{
							"name": 12,
						},
					},
				},
				path:  "hello.user.name",
				value: 42,
			},
			want: map[string]interface{}{
				"hello": map[interface{}]interface{}{
					"user": map[interface{}]interface{}{
						"name": 42,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test5",
			args: args{
				original: map[string]interface{}{
					"hello": map[interface{}]interface{}{
						"user": map[interface{}]interface{}{
							"name": 12,
						},
					},
				},
				path:  "hello",
				value: "bar",
			},
			want: map[string]interface{}{
				"hello": "bar",
			},
			wantErr: false,
		},
		{
			name: "test6",
			args: args{
				original: map[string]interface{}{
					"$goodbye": "john",
					"$hello": map[interface{}]interface{}{
						"john": map[interface{}]interface{}{
							"age": 12,
						},
					},
				},
				path:  "$hello[$goodbye]age",
				value: 66,
			},
			want: map[string]interface{}{
				"$goodbye": "john",
				"$hello": map[interface{}]interface{}{
					"john": map[interface{}]interface{}{
						"age": 66,
					},
				},
			},
			wantErr: false,
		},
		{
			name: "test7",
			args: args{
				original: map[string]interface{}{
					"$goodbye": "john",
					"$hello": map[interface{}]interface{}{
						"john": map[interface{}]interface{}{
							"children": []interface{}{
								"alice", "bob", "cecil",
							},
						},
					},
				},
				path:  "$hello[$goodbye]children[0]",
				value: "YOLO",
			},
			want: map[string]interface{}{
				"$goodbye": "john",
				"$hello": map[interface{}]interface{}{
					"john": map[interface{}]interface{}{
						"children": []interface{}{
							"YOLO", "bob", "cecil",
						},
					},
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := testAddPathHelper(tt.args.original, tt.args.path, tt.args.value)
			if (err != nil) != tt.wantErr {
				t.Errorf("test_AddPath_Helper() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("test_AddPath_Helper() = %v, want %v", got, tt.want)
			}
		})
	}
}
