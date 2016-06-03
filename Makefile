all:
	make -C internal -j4

config:
	cd internal && ./autogen.sh \
		--with-jemalloc-prefix="je_" --disable-valgrind

clean distclean:
	make -C internal $@
