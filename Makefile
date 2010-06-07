source := global.go lexer.go
objects := $(patsubst %.go,%.8,$(source))

.PHONY: clean

main: $(objects)

%.8 : %.go
	8g -o $@ $<

% : %.8
	8l -o $@ $<

clean:
	rm -v -f main $(objects)
