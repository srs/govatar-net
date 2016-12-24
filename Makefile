PLATFORMS := linux/amd64 windows/amd64/.exe windows/386/.exe darwin/amd64
VERSION = 0.1.0

TEMP = $(subst /, ,$@)
OS = $(word 1, $(TEMP))
ARCH = $(word 2, $(TEMP))
EXT = $(word 3, $(TEMP))
LDFLAGS = "-X main.Version=${VERSION}"

DOCKER_OWNER = stenrs
DOCKER_IMAGE = govatar-net
DOCKER_QNAME = $(DOCKER_OWNER)/$(DOCKER_IMAGE)
DOCKER_BUILD_TAG = $(DOCKER_QNAME):$(VERSION)
DOCKER_LATEST_TAG = $(DOCKER_QNAME):latest

build: clean $(PLATFORMS)

clean:
	rm -rf build/

$(PLATFORMS):
	GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags $(LDFLAGS) -o 'build/$(OS)-$(ARCH)/govatar-net$(EXT)' github.com/srs/govatar-net
	zip 'build/$(OS)-$(ARCH)/govatar-net.zip' 'build/$(OS)-$(ARCH)/govatar-net$(EXT)'

docker-build: build
	docker build -t $(DOCKER_BUILD_TAG) .
	docker tag $(DOCKER_BUILD_TAG) $(DOCKER_LATEST_TAG)

docker-login:
	docker login -u "$(DOCKER_OWNER)" -p "$(DOCKER_PASS)"

docker-push: docker-build docker-login
	docker push $(DOCKER_BUILD_TAG)
 	docker push $(DOCKER_LATEST_TAG)
 
