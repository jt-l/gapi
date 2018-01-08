BINARY=TT

VERSION=1.0.0
BUILD=`git rev-parse HEAD`

LDFLAGS=-ldflags "-X main.Version=${VERSION} -X main.Build=${BUILD}"

.DEFAULT_GOAL: ${BINARY}

${BINARY}:
	go build ${LDFLAGS} -o ${BINARY} cmd/TT/main.go

install: 
	go install ${LDFLAGS} -o ${BINARY} cmd/TT/main.go

test:
	go test

clean: 
	if [ -f ${BINARY} ] ; then rm ${BINARY} ; fi

.PHONY: clean install
