// +build windows
// +build 386 amd64

package internal

/*
#include <stdint.h>

typedef struct uint64aligned { uint64_t value[1]; } uint64aligned;
typedef struct int64aligned { int64_t value[1]; } int64aligned;
*/
import "C"

func GetUint64(u64 C.uint64aligned) uint64 {
	return uint64(u64.value[0])
}

func SetUint64(u64 *C.uint64aligned, value uint64) {
	u64.value[0] = C.uint64_t(value)
}

func GetInt64(i64 C.int64aligned) int64 {
	return int64(i64.value[0])
}

func SetInt64(i64 *C.int64aligned, value int64) {
	i64.value[0] = C.int64_t(value)
}