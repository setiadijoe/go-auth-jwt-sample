# Golang Package Struct

This package is just for Go package struct. On this example, I wrote for using JWT and GOKit without accessing any thirdparty. Here it is the structure of this package

```sh
-cmd 
    ->main.go
-endpoint
    ->endpoint.go
    ->request.go
    ->response.go
-model
    ->auth_model.go
-repository
    ->repository_interface.go
--transport
    -http
    -grpc
-usecase
    ->usecase_interface.go
go.mod
go.sum
```