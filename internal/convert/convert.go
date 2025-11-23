// Comment: C与Go之间的转换

package convert

import (
	"strings"
	"unsafe"
)

// 安全地将Go字符串转换为C风格的空终止字节数组，并返回指向第一个字节的指针
func ToBytePtr(s string) *byte {
	size := len(s) + 1
	if index := strings.IndexByte(s, 0); index != -1 {
		size = index + 1
	}

	result := make([]byte, size)
	copy(result, s)
	return &result[0]
}

// 同上，带空值检查，空字符串返回nil
func ToBytePtrNullable(s string) *byte {
	if len(s) == 0 {
		return nil
	}
	return ToBytePtr(s)
}

// 将C风格的空终止字符串指针转换回Go字符串
func ToString(p *byte) string {
	if p == nil {
		return ""
	}
	i := 0
	for ptr := unsafe.Pointer(p); *(*byte)(unsafe.Add(ptr, i)) != 0; i++ {
	}
	return string(unsafe.Slice(p, i))
}

// 空终止的字符串指针数组转换为Go的字符串切片，char *[]/char ** -> []string
func ToStringSlice(pointers **byte) []string {
	if pointers == nil {
		return nil
	}

	strings := make([]string, 0)

	for ptr := unsafe.Pointer(pointers); *(**byte)(ptr) != nil; ptr = unsafe.Add(ptr, 8) {
		strings = append(strings, ToString(*(**byte)(ptr)))
	}

	return strings
}
