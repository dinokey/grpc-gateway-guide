module github.com/example/server/gateway

go 1.20

replace github.com/example/service/employee => ../employee-service/example/demo/v1/service

replace github.com/example/types => ../type/example/demo/v1/type

require (
	github.com/example/service/employee v0.0.0-00010101000000-000000000000
	github.com/example/service/salary v0.0.0-00010101000000-000000000000
	github.com/golang/glog v1.0.0
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.15.0
	google.golang.org/grpc v1.53.0
)

require (
	github.com/example/types v0.0.0-00010101000000-000000000000 // indirect
	github.com/golang/protobuf v1.5.2 // indirect
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20230209215440-0dfe4f8abfcc // indirect
	google.golang.org/protobuf v1.28.1 // indirect
)

replace github.com/example/service/salary => ../salary-service/example/demo/v1/service

replace github.com/samsung/magic/v1/api/userService => ../../test/grpc_user_service

replace github.com/samsung/magic/v1/type => ../../test/grpc_user_type
