package js

import (
	"github.com/xymeng16/ohos-go"
)

type Value struct {
	Env   Env
	Value napi.Value
}

func (v Value) GetEnv() Env {
	return v.Env
}
