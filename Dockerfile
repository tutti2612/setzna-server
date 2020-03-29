FROM golang:1.14

WORKDIR /go/src/setzna
RUN go get github.com/derekparker/delve/cmd/dlv && \
    go get github.com/oxequa/realize && \
    go build -o /go/bin/dlv github.com/go-delve/delve/cmd/dlv