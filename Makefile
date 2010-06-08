include $(GOROOT)/src/Make.$(GOARCH)

all: main

DIRS=\
	lexer\
	parser\

.PHONY: dirs $(DIRS)

dirs: $(DIRS)

$(DIRS):
	$(MAKE) -C $@ install

parser: lexer

main: dirs

TARG=main
GOFILES=\
	main.go\

include $(GOROOT)/src/Make.cmd
