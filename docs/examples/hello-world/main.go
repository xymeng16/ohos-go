package main

import (
	"fmt"

	"github.com/xymeng16/ohos-go"
	"github.com/xymeng16/ohos-go/entry"
)

func init() {
	entry.Export("hello", HelloHandler)
}

func HelloHandler(env napi.Env, info napi.CallbackInfo) napi.Value {
	fmt.Println("hello world!")
	return nil
}

func main() {}
