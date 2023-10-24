package main

import (
	"fmt"
	jsoniter "github.com/go-shafaq/jsoniter"
	"strings"
)

func main() {
	jsoniter.DefCase(func(tag string) func(pkgPath string, name string) string {
		return func(pkgPath, name string) string {
			println(pkgPath, name)
			return strings.ToUpper(name)
		}
	})
	marshal, err := jsoniter.Marshal(u)

	fmt.Println(string(marshal), err)
}

var u = &User{13, 56}

type User struct {
	Id       int
	ParentId int
}
