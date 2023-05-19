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

# generate hertz code
cd hz-server
hz new -idl ../idl/student_management.thrift -module hz-server


```