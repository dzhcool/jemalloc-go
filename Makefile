.DEFAULT_GOAL = build-all

PWD := $(shell pwd)
SRC := jemalloc-4.4.0

build-all:
	@test -f $(SRC)/Makefile || make config --quiet
	@make -C $(SRC) install_lib_static -j8

config:
	cd $(SRC) && ./autogen.sh --with-jemalloc-prefix="je_" --libdir=$(PWD)/install

clean distclean:
	@test -f $(SRC)/Makefile && make -C $(SRC) --quiet distclean || true

install: build-all
	go install ./

test:
	go test -v ./
