# Build file for bbelial

PROGNAME =	bbelial

MAINFILE =	main.go
BUILDDIR =	build

all: build

templ:
	templ generate

run: templ
	go run ${MAINFILE}

build: templ
	@mkdir -pv ${BUILDDIR}
	go build -o ${BUILDDIR}/${PROGNAME} bbelial

clean:
	@rm -rfv ${BUILDDIR}

PHONY: build run templ
