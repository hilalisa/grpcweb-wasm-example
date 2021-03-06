generate:
	protoc -I. ./proto/web.proto \
		--go_out=plugins=grpc:$$GOPATH/src
	go1.11rc1 generate -x ./frontend/

serve:
	go1.11rc1 run main.go
