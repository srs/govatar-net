PLATFORMS := linux/amd64 windows/amd64/.exe windows/386/.exe darwin/amd64
VERSION = '0.1.0'

TEMP = $(subst /, ,$@)
OS = $(word 1, $(TEMP))
ARCH = $(word 2, $(TEMP))
EXT = $(word 3, $(TEMP))
LDFLAGS = "-X main.Version=${VERSION}"

build: clean $(PLATFORMS)

clean:
	rm -rf build/

$(PLATFORMS):
	GOOS=$(OS) GOARCH=$(ARCH) go build -ldflags $(LDFLAGS) -o 'build/govatar-net$(EXT)' github.com/srs/govatar-net
	zip 'build/govatar-net-$(OS)-$(ARCH)-$(VERSION).zip' 'build/govatar-net$(EXT)'
