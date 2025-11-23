// Comment: C与Go之间的内存转换

package mem

import "unsafe"

// 将C数组复制到Go切片中
func Copy[T any, N ~int32](array *T, count N) []T {
	if array == nil {
		return nil
	}
	result := make([]T, count)
	copy(result, unsafe.Slice(array, count))
	return result
}

// 复制每个C指针所指向的底层数据到Go切片中
func DeepCopy[T any, N ~int32](array **T, count N) []*T {
	if array == nil {
		return nil
	}
	dst := make([]*T, count)
	src := unsafe.Slice(array, count)
	for i := 0; i < int(count); i++ {
		if src[i] != nil {
			value := *src[i]
			dst[i] = &value
		}
	}
	return dst
}
