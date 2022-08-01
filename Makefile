BINARY_PATH="./build"
BINARY_NAME=""
PLATFORM=windows # windows / darwin / linux

build:
	go build -o ${BINARY_PATH}/ ./...
#	GOARCH=amd64 GOOS=${PLATFORM} go build -o ${BINARY_PATH}/ ./...

# ifeq ($(BINARY_NAME),"")
# 	GOARCH=amd64 GOOS=${PLATFORM} go build -o ${BINARY_PATH}/ ./...
# else
# 	GOARCH=amd64 GOOS=${PLATFORM} go build -o ${BINARY_PATH}/${BINARY_NAME} cmd/leasing-car/main.go
# endif

run:
	${BINARY_PATH}/${BINARY_NAME}

build_and_run: build run

clean:
	go clean
#	rm ${BINARY_PATH}/${BINARY_NAME}-windows
