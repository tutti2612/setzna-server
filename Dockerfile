FROM golang:1.14

ENV SRC_DIR=/go/src/setzna
WORKDIR $SRC_DIR
RUN go get github.com/pilu/fresh
CMD ["fresh"]