v ?= latest

GOVERSION=`go version`
BINARYVERSION=`git tag`
GITLASTLOG=$(shell git log --pretty=medium -1)
LDFLAGS=-ldflags '-w -X "main.BinaryVersion=${BINARYVERSION}" -X "main.GoVersion=${GOVERSION}" -X "main.GitLastLog=${GITLASTLOG}"'

.PHONY: build
build: 
	env CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build ${LDFLAGS} -o combat main.go

.PHONY: build_win
build_win: 
	env CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build ${LDFLAGS} -o combat.exe main.go

.PHONY: docker
docker: build
	docker build . -t combat:latest

.PHONY: test
test:
	go test -v ./... -cover

.PHONY: push
push:
	# docker hub
	docker login -u hellodudu86 -p hello1986
	docker tag combat hellodudu86/combat:$(v)
	docker push hellodudu86/combat:$(v)

.PHONY: push_coding
push_coding:
	# coding hub
	docker tag combat mmstudio-docker.pkg.coding.net/blade/server/combat:$(v)
	docker push mmstudio-docker.pkg.coding.net/blade/server/combat:$(v)

	
.PHONY: push_github
push_github:
	# github packages
	docker tag combat ghcr.io/east-eden/server/combat:$(v)
	docker push ghcr.io/east-eden/server/combat:$(v)

.PHONY: clean
clean:
	make -C ../../ clean
