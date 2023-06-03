deps:
	go mod download && go mod tidy

carecompanion:
	go build -o $@ ./cmd/

autogen:
	protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative ./proto/*.proto


