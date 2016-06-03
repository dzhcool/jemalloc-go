// Package jemalloc uses the cgo compilation facilities to build the
// jemalloc library.
package jemalloc

// #cgo CFLAGS: -Iinternal/include -std=gnu99
// #cgo CPPFLAGS: -D_REENTRANT
// #cgo linux CPPFLAGS: -D_GNU_SOURCE
// #cgo LDFLAGS: -lm
import "C"
