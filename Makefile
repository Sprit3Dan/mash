PROTOC = protoc
GO = go

all: proto
	cd ./cli; ${GO} build -o ./../bin/mash main.go

proto:
	${PROTOC} --go_out=. ./**/*.proto

clean:
	rm ./cli/main; find . -name "*.pb.go" -delete
