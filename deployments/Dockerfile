FROM iron/go:dev

#ENV PATH=/Users/james/go/src/github.com/james/TT

WORKDIR /app

ADD . /go/src/github.com/james/TT 

ENV PORT="8090"

RUN cd /go/src/github.com/james/TT/cmd/TT; go build  ; cp TT /app/

ENTRYPOINT ["./TT"]
