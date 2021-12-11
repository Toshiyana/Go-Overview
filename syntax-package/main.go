package main

import (
	"fmt"

	"go-mod-practice/helpers"
)

func main() {

	// How to create my package
	// 1. pod mod init [package-name]
	// package-name: localで作ったpackageを用いる場合、任意の名前で良い。remoteのパッケージを用いる場合、urlなどを指定。
	// 今回はlocalで作ったpackageのhelpersを利用。

	// 2.
	// localの場合： go build
	// remoteの場合：go mod tidy（多分）

	var myVar helpers.SomeType
	myVar.TypeName = "Some Name"
	fmt.Println(myVar.TypeName)
}
