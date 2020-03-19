FROM golang:1.14

ENV SRC_DIR=/go/src/setzna
WORKDIR $SRC_DIR

# ADD . $SRC_DIR