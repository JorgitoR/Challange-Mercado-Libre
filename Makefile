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


migrateup:
	migrate -path internal/infraestructure/db/migration -database "postgresql://root:MaXRn0aWBcFEnmPlmuzC@database-1.ctmmrijpqxtv.us-east-2.rds.amazonaws.com:5432/mercado_libre"  force 15 up

	

migratedown:
	migrate -path internal/infraestructure/db/migration -database "postgresql://root:MaXRn0aWBcFEnmPlmuzC@database-1.ctmmrijpqxtv.us-east-2.rds.amazonaws.com:5432/mercado_libre" -verbose down