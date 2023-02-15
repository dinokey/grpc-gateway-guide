module github.com/example/service/employee

go 1.20

replace github.com/example/types => ../../../../../type/example/demo/v1/type

require (
	github.com/example/types v0.0.0-00010101000000-000000000000
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0
	google.golang.org/genproto v0.0.0-20230209215440-0dfe4f8abfcc
	google.golang.org/grpc v1.53.0
	google.golang.org/protobuf v1.28.1
)

require (
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
)
