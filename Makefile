run:
	go run cmd/main.go

deps:
	goproxy=direct \
	GOSUMDB=off \
	go mod tidy

test:
	go test ./...  -cover -coverprofile=coverage.out 

gen-mocks:
	mockery --dir=./internal --all 

# formats project with go's style guidelines
fmt:
	gofmt -w -s ./internal ./cmd

watch-coverage:
	go tool cover -html=coverage.out
