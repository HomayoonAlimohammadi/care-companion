.PHONY: deps carecompanion autogen 

deps:
	go mod download && go mod tidy

carecompanion:
	go build -o $@ ./cmd/

autogen:
	protoc --proto_path=proto \
	--go_out=. \
	--go_opt=Mcare_companion.proto=autogen/care_companion \
	--go-grpc_opt=Mcare_companion.proto=autogen/care_companion \
	--go-grpc_out=. \
	care_companion.proto
	