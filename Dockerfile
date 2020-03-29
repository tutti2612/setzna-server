FROM golang:1.14

WORKDIR /go/src/setzna
RUN go get github.com/derekparker/delve/cmd/dlv && \
    go get github.com/oxequa/realize

#CMD [ "realize", "start", "--run" ]
#CMD ["dlv", "debug", "--headless", "--listen=:2345", "--api-version=2"]