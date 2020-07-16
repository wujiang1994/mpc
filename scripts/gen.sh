cd testapp/app/protos
protoc mpc.proto --go_out=plugins=grpc:../pbs
# add -I need specify an exact prefix, if not, use current directory