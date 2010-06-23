include $(GOROOT)/src/Make.$(GOARCH)

all: packages main

.PHONY: all

PACKAGES=\
	lexer\
	parser\

.PHONY: packages $(PACKAGES)

packages: $(PACKAGES)

$(PACKAGES):
	$(MAKE) -C $@ install

parser: lexer

main: packages

TARG=main
GOFILES=\
	main.go\

include $(GOROOT)/src/Make.cmd
