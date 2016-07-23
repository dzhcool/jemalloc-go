.DEFAULT_GOAL = all

config:
	cd jemalloc && ./autogen.sh --with-jemalloc-prefix="je_"
	@make -f helps.mk --quiet relink

ifeq ("$(wildcard jemalloc/Makefile)","")

all: config

clean distclean:

else

all:

clean distclean:
	@make -C jemalloc --quiet distclean
	@make -f helps.mk --quiet rmlink

endif
