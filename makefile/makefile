gen:
	protoc --proto_path=../protos ../protos/proto/*.proto --go_out=plugins=grpc:../protos/pb
clean:
	rm -R ../protos/pb/*
run:
	go run ../app/main.go