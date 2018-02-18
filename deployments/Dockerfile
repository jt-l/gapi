FROM iron/go:dev

WORKDIR /app

ADD . /go/src/github.com/james/TT 

ENV PORT="8090"

RUN cd /go/src/github.com/james/TT/cmd/TT; go build  ; cp TT /app/

ENTRYPOINT ["./TT"]
