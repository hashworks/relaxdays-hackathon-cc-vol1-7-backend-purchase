run: *.go swag
	CGO_ENABLED=1 go run main.go

deps:
	go get -u github.com/swaggo/swag/cmd/swag

swag: */*.go
	swag init


