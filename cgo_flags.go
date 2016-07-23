package jemalloc

// #cgo CFLAGS: -Ijemalloc/include -std=gnu99
// #cgo       CPPFLAGS: -D_REENTRANT
// #cgo linux CPPFLAGS: -D_GNU_SOURCE
// #cgo LDFLAGS: -lm
import "C"
