# activity-log

Another grpc example

go mod init github.com/CodingForFunAndProfit/activity-log

go get google.golang.org/protobuf/cmd/protoc-gen-go
go get google.golang.org/grpc/cmd/protoc-gen-go-grpc

protoc api/v1/activityproto \
 --go_out=. \
 --go_opt=paths=source_relative \
 --proto_path=.

CTRL + SHIFT + P: reload window

protoc api/v1/activity.proto \
 --go-grpc_out=. \
 --go-grpc_opt=paths=source_relative \
 --proto_path=.

go mod tidy

CTRL + SHIFT + P: reload window

Or:
protoc api/v1/\*.proto \
 --go_out=. \
 --go_opt=paths=source_relative \
 --go-grpc_out=. \
 --go-grpc_opt=paths=source_relative \
 --proto_path=.
