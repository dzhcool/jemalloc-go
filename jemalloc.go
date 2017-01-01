package jemalloc

// #cgo         CPPFLAGS: -I.
// #cgo          LDFLAGS: -Llib -ljemalloc
// #cgo  darwin  LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
// #cgo !darwin  LDFLAGS: -Wl,-unresolved-symbols=ignore-all
// #include <jemalloc/jemalloc.h>
import "C"

import "unsafe"

func Calloc(count, size int) unsafe.Pointer {
	return C.je_calloc(C.size_t(count), C.size_t(size))
}

func Malloc(size int) unsafe.Pointer {
	return C.je_malloc(C.size_t(size))
}

func Valloc(size int) unsafe.Pointer {
	return C.je_valloc(C.size_t(size))
}

func Realloc(ptr unsafe.Pointer, size int) unsafe.Pointer {
	return C.je_realloc(ptr, C.size_t(size))
}

func Free(ptr unsafe.Pointer) {
	C.je_free(ptr)
}
