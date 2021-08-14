
default:
	GOOS=linux GOARCH=arm GOARM=7 go build -o lumagen main.go

test:
	go test ./...
