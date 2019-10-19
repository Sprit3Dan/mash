PROTOC = protoc
GO = go

all: proto
	cd ./cli; ${GO} build main.go

proto:
	${PROTOC} --go_out=. ./**/*.proto
