before.build:
	go mod download && go mod vendor

build.wslight:
	@echo "build in ${PWD}";env GOOS=windows GOARCH=amd64 go build -o wslight.exe cmd/wslight/main.go
