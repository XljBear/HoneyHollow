.PHONY : mac windows linux all mkdir run web
mac: prepare
	 CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/HoneyHollowServer_Mac/HoneyHollowServer main.go

windows: prepare
	go generate
	CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/HoneyHollowServer_Windows/HoneyHollowServer.exe ./

linux: prepare
	CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -ldflags "-s -w" -trimpath -o ./build/HoneyHollowServer_Linux/HoneyHollowServer main.go
run:
	go run main.go
all: mac windows linux

web:
	cd frontend && yarn build
prepare:
	mkdir -p ./build/
	mkdir -p ./build/HoneyHollowServer_Mac
	mkdir -p ./build/HoneyHollowServer_Windows
	mkdir -p ./build/HoneyHollowServer_Linux