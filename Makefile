# all: get-deps fmt darwin linux arm windows build coverage
all: clean create build

clean:
	rm -rf binaries

create:
	mkdir binaries
	mkdir binaries/linux_386 binaries/linux_amd64 binaries/windows_amd64 binaries/windows_386 binaries/darwin_arm64 binaries/darwin_amd64

build:
	GOARCH=386 && GOOS=linux go build -o binaries/linux_386/dm-cmd main.go
	GOARCH=amd64 && GOOS=linux go build -o binaries/linux_amd64/dm-cmd main.go

	GOARCH=386 && GOOS=windows go build -o binaries/windows_386/dm-cmd.exe main.go
	GOARCH=amd64 && GOOS=windows go build -o binaries/windows_amd64/dm-cmd.exe main.go

	GOARCH=amd64 && GOOS=darwin go build -o binaries/darwin_amd64/dm-cmd main.go
	GOARCH=arm64 && GOOS=darwin go build -o binaries/darwin_arm64/dm-cmd main.go

	






