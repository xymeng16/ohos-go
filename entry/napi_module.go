package entry

/*
#include <stdlib.h>

#include "./entry.h"
*/
import "C"

import (
	"github.com/xymeng16/ohos-go"
)

//export InitializeModule
func InitializeModule(cEnv C.napi_env, cExports C.napi_value) C.napi_value {
	env, exports := napi.Env(cEnv), napi.Value(cExports)
	napi.InitializeInstanceData(env)

	for _, export := range napiGoGlobalExports {
		cb, _ := napi.CreateFunction(env, export.Name, export.Callback)
		name, _ := napi.CreateStringUtf8(env, export.Name)
		napi.SetProperty(env, exports, name, cb)
	}

	return cExports
}
