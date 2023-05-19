# HTTP Server(Hertz) <-> RPC Server(Kitex)

```shell
# install hertz
go install github.com/cloudwego/hertz/cmd/hz@latest
# install kitex
go install github.com/cloudwego/kitex/tool/cmd/kitex@latest

# create directory
mkdir -p hz_kx_demo
cd hz_kx_demo
mkdir -p idl
mkdir -p hz-server
mkdir -p kx-server
# create idl

# generate hertz project
cd hz-server
hz new -idl ../idl/student_management.thrift -module hz-server
# generate rpc client
kitex -module hz-server ../idl/student_management.thrift


# generate kitex project
cd kx-server
kitex -module kx-server -service kx-server ../idl/student_management.thrift

```


Reference: 
- https://github.com/cloudwego/hertz-examples/tree/main/hz_kitex_demo