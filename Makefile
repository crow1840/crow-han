DOCKER_IMAGE_TAG=latest

build:
	mkdir -p bin
	cd bin && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -installsuffix cgo  crow-han/cmd/auth
	cd bin && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -installsuffix cgo  crow-han/cmd/user
	cd bin && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -mod vendor -a -installsuffix cgo  crow-han/cmd/gate-way

swag:
	cd internal/app/gate-way/service
	swag init --parseDependency --parseInternal --parseDepth 2 -g .\router.go

docker-build: build
	docker build -f docker/Dockerfile_auth . -t crow/han-auth:$(DOCKER_IMAGE_TAG)
	docker build -f docker/Dockerfile_user . -t crow/han-user:$(DOCKER_IMAGE_TAG)
	docker build -f docker/Dockerfile_gateway . -t crow/han-gateway:$(DOCKER_IMAGE_TAG)