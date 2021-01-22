package main

import (
	"fmt"
	"github.com/gogf/gf/encoding/gjson"
	"github.com/gogf/gf/frame/g"
	"github.com/wowqhb/goframe_demo/dao"
)

func main() {
	g.Cfg().SetPath("./")
	g.Cfg().SetFileName("config.yaml")

	all, err := dao.Book.FindAll()
	if err != nil {
		panic(err)
	}
	for i, book := range all {
		encode, err := gjson.Encode(book)
		if err != nil {
			panic(err)
		}
		fmt.Printf("index: %d, val: %s\n", i, string(encode))
	}

	one, err := dao.Book.Limit(1).FindOne(g.Map{"book": "golang高级编程", "price": 59})
	if err != nil {
		panic(err)
	}
	encode, err := gjson.Encode(one)
	if err != nil {
		panic(err)
	}
	fmt.Println(string(encode))
}
