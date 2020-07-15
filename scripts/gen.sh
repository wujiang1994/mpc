cd testapp/app/protos
protoc mpc.proto --go_out=plugins=grpc:../pbs2
# add -I need specify an exact prefix, if not, use current directory