package napi

/*
#cgo CFLAGS: -DDEBUG
#cgo CFLAGS: -D_DEBUG
#cgo CFLAGS: -DV8_ENABLE_CHECKS
#cgo CFLAGS: -DNAPI_EXPERIMENTAL
#cgo CFLAGS: -I$OHOS_NDK_HOME/native/sysroot/usr/include
#cgo CFLAGS: -Wno-error

#cgo darwin LDFLAGS: -Wl,-undefined,dynamic_lookup
#cgo darwin LDFLAGS: -Wl,-no_pie
#cgo darwin LDFLAGS: -Wl,-search_paths_first
#cgo darwin LDFLAGS: -arch x86_64

#cgo linux LDFLAGS: -Wl,-unresolved-symbols=ignore-all

#cgo LDFLAGS: -L$OHOS_NDK_HOME/native/sysroot/usr/lib/aarch64-linux-ohos
*/
import "C"
