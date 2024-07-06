PROG=bin/WarmStorageChaind

SRCS=./cmd/WarmStorageChaind

INSTALL_PREFIX=/usr/local/WarmStorageChaind

CONF_INSTALL_PREFIX=/usr/local/WarmStorageChaind

# git commit hash
COMMIT_HASH=$(shell git rev-parse --short HEAD || echo "GitNotFound")

# build time
BUILD_DATE=$(shell date '+%Y-%m-%d %H:%M:%S')

# build opts
CFLAGS = -ldflags "-s -w  -X \"main.BuildVersion=${COMMIT_HASH}\" -X \"main.BuildDate=$(BUILD_DATE)\""

all: build

build:
	if [ ! -d "./bin/" ]; then \
	mkdir bin; \
	fi
	go build $(CFLAGS) -o $(PROG) $(SRCS)

# build with race
race:
	if [ ! -d "./bin/" ]; then \
    	mkdir bin; \
    	fi
	go build $(CFLAGS) -race -o $(PROG) $(SRCS)