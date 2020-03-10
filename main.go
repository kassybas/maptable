package main

import (
	"fmt"

	"github.com/kassybas/vartable/vartable"
	"github.com/kataras/golog"
)

func main() {
	vt := vartable.New()
	orig := map[interface{}]interface{}{
		"user": map[interface{}]interface{}{
			"name": "john",
		},
	}
	err := vt.AddPath("$hello", orig)
	if err != nil {
		golog.Error(err)
	}
	fmt.Println(vt)
	err = vt.AddPath("$hello.user.yolo", "JANE")
	if err != nil {
		golog.Error(err)
	}
	fmt.Println(vt)
}
