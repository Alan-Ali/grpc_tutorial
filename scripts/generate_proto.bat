@echo off
setlocal

REM Set the path to the protoc binary
set PROTOC_PATH=C:\protoc-25.4-win64\bin

REM Set the path to the proto files
set PROTO_FILES_PATH=..\protobufs

REM Set the output path for the generated files
set OUT_PATH=..

REM Run protoc with the specified paths
%PROTOC_PATH%\protoc --proto_path=%PROTO_FILES_PATH% --go_out=%OUT_PATH% --go-grpc_out=%OUT_PATH% %PROTO_FILES_PATH%\backend_gateway.proto

endlocal
