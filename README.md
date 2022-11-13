```
protoc --go_out=./pbs \
         --go_opt=paths=source_relative \
         --go-grpc_out=./pbs \
         --go-grpc_opt=paths=source_relative \
         --proto_path=protos \
         protos/*.proto
```

```
go run ./server/echo
```

```
go run ./client/echo
```
