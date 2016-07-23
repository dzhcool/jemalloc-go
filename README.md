# jemalloc

``bash
$ mkdir -p $GOPATH/src/github.com/spinlock && cd $_
$ git clone git@github.com:spinlock/jemalloc-go.git
$ make install
```

```go
package demo

// #cgo         CPPFLAGS: -I <relative-path>/jemalloc-go/jemalloc/include
// #cgo  darwin LDFLAGS: -Wl,-undefined -Wl,dynamic_lookup
// #cgo !darwin LDFLAGS: -Wl,-unresolved-symbols=ignore-all
// #include <jemalloc/jemalloc.h>
import "C"
import _ "github.com/spinlock/jemalloc-go"

func malloc(n int) unsafe.Pointer {
    return C.je_malloc(C.size_t(n))
}
```
