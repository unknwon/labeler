.PHONY: install release

BIN_NAME=labeler

all: install

install:
	go install -v

release:
	env GOOS=darwin GOARCH=amd64 go build -o ${BIN_NAME}; tar czf darwin_amd64.tar.gz ${BIN_NAME}
	env GOOS=linux GOARCH=amd64 go build -o ${BIN_NAME}; tar czf linux_amd64.tar.gz ${BIN_NAME}
	env GOOS=linux GOARCH=386 go build -o ${BIN_NAME}; tar czf linux_386.tar.gz ${BIN_NAME}
	env GOOS=linux GOARCH=arm go build -o ${BIN_NAME}; tar czf linux_arm.tar.gz ${BIN_NAME}
	env GOOS=windows GOARCH=amd64 go build -o ${BIN_NAME}; tar czf windows_amd64.tar.gz ${BIN_NAME}
	env GOOS=windows GOARCH=386 go build -o ${BIN_NAME}; tar czf windows_386.tar.gz ${BIN_NAME}
