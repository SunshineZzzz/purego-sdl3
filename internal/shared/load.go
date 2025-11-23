// Comment: 加载动态链接库并获取符号地址，主要用于非Windows系统

//go:build !windows

package shared

import "github.com/ebitengine/purego"

func Load(name string) (uintptr, error) {
	return purego.Dlopen(name, purego.RTLD_LAZY)
}

func Get(lib uintptr, name string) uintptr {
	addr, err := purego.Dlsym(lib, name)
	if err != nil {
		panic(err)
	}
	return addr
}
