all:

config:
	cd internal && ./autogen.sh \
		--with-jemalloc-prefix="je_" --disable-valgrind

build:
	make -C internal -j4
	go build

clean distclean:
	make -C internal $@
