#
.PHONY: build test coverage run clean

BINARY_NAME=string-reduce
BINARY_FOLDER=binaries
MODULE_NAME=github.com/gavinmathias/string-reduce
BUILD_VERSION=$(shell cat VERSION.txt)
BUILD_TIME=$(shell date)
SEPARATOR="+---------------------------------+"

all: build test run

build:
	@figlet build
	@printf "Making the binary folder ${BINARY_FOLDER} if it doesn't exist\n"
	mkdir -p ${BINARY_FOLDER}
	@printf "${SEPARATOR}\n"

	@printf "Downloading external vendor modules if necessary\n"
	go mod vendor
	@printf "${SEPARATOR}\n"
	
	@printf "Compiling for MacOS\n"
	GOARCH=amd64 GOOS=darwin go build -o ${BINARY_FOLDER}/${BINARY_NAME} \
	 	-ldflags \
		"-X '${MODULE_NAME}/build.Version=${BUILD_VERSION}' \
		-X '${MODULE_NAME}/build.Time=${BUILD_TIME}'" \
		cmd/main.go
	@printf "${SEPARATOR}\n"

run:
	@figlet run
	./${BINARY_FOLDER}/${BINARY_NAME} -version
	@printf "${SEPARATOR}\n\n"
	./${BINARY_FOLDER}/${BINARY_NAME} -consecutive 3 -input 112233444321
	@printf "${SEPARATOR}\n\n"
	./${BINARY_FOLDER}/${BINARY_NAME} -consecutive 3 -input 11223344431
	@printf "${SEPARATOR}\n\n"

test:
	@figlet test
	go test -v -cover -tags test -coverprofile=coverage.txt ./...

coverage: test
	@figlet coverage
	go tool cover -html=coverage.txt

clean:
	@figlet clean
	go clean
	@printf "${SEPARATOR}\n"
	rm -f ${BINARY_FOLDER}/${BINARY_NAME}*
	@printf "${SEPARATOR}\n"