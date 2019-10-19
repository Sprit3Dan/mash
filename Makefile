PROTOC = protoc
GO = go

all: proto
	cd ./src/cli; ${GO} build -o ../../bin/mash main.go

proto:
	${PROTOC} --go_out=. ./src/packages/**/*.proto

clean:
	rm ./cli/main; find . -name "*.pb.go" -delete
