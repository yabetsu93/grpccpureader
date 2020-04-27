gen:
	protoc --proto_path=../protos ../protos/proto/*/*.proto --go_out=plugins=grpc:../protos/pb
clean:
	rm ../protos/pb/*/*.go
run:
	go run ../app/main.go