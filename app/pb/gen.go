package pb

//go:generate protoc -I. -I./src --swagger_out=allow_delete_body=true:../../swagger --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative serve.proto
//go:generate  protoc -I. -I./src --grpc-gateway_out . --grpc-gateway_opt logtostderr=true --grpc-gateway_opt paths=source_relative serve.proto
