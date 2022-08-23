BINARY_NAME=pastebin

build: clean
	cd frontend && yarn build && cd ..
	go build -o build/${BINARY_NAME}-darwin
#	GOARCH=amd64 GOOS=darwin go build -o build/${BINARY_NAME}-darwin main.go
#	GOARCH=amd64 GOOS=linux go build -o build/${BINARY_NAME}-linux main.go
#	GOARCH=amd64 GOOS=window go build -o build/${BINARY_NAME}-windows main.go

clean:
	go clean
	rm -rf frontend/build/
	rm -rf build/
#	rm -f ${BINARY_NAME}-darwin
#	rm -f ${BINARY_NAME}-linux
#	rm -f ${BINARY_NAME}-windows
