# multiplica

## INSTALL ##

Download the code to $GOPATH/src. The source code includes the already compiled binaries an a unit test for server_grpc:

* server_grpc
* server_grpc_test.go
* client_grpc
* server_http

The source code can be compiled running:

```
go build $GOPATH/multiplica/multiplica_server/server_grpc.go
go build $GOPATH/multiplica/multiplica_client/client_grpc.go
go build $GOPATH/multiplica/multiplica_http/server_http.go
```

## gRPC ##
For gRPC the server listens in port 9500. Run the server_grpc and then open other terminal and run the client
```
$GOPATH/multiplica/multiplica_client/client_grpc 5 3
```

The server will return the multiplication of 5*3 = 15

Note: protocol buffer already compiled are included with the souce code.

## HTTP ##
Run the http server which will listen in port 8080:

```
$GOPATH/multiplica/multiplica_http/server_http
```

Then from another terminal query the http server:

```
curl "http://localhost:8080/multiplica/3/5"
```

The http server will answer with the product of 3*5 = 15
