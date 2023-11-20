// +build windows
// +build 386 amd64

package internal

/*
#include <stdint.h>

typedef struct uint64aligned { uint64_t value[1]; } uint64aligned;
typedef struct int64aligned { int64_t value[1]; } int64aligned;
*/
import "C"

type Uint64Aligned struct {
    value C.uint64aligned
}

type Int64Aligned struct {
    value C.int64aligned
}

func (u64 *Uint64Aligned) Get() uint64 {
    return uint64(u64.value.value[0])
}

func (u64 *Uint64Aligned) Set(value uint64) {
    u64.value.value[0] = C.uint64_t(value)
}

func (i64 *Int64Aligned) Get() int64 {
    return int64(i64.value.value[0])
}

func (i64 *Int64Aligned) Set(value int64) {
    i64.value.value[0] = C.int64_t(value)
}