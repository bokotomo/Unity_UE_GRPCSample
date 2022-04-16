FROM golang:1.18.1-stretch

SHELL ["/bin/bash", "-c"]
RUN apt update && apt-get install -y vim unzip

# install protc
WORKDIR /protoc
RUN wget https://github.com/protocolbuffers/protobuf/releases/download/v3.15.8/protoc-3.15.8-linux-x86_64.zip
RUN unzip protoc-3.15.8-linux-x86_64.zip
RUN ln -s /protoc/bin/protoc /bin/protoc

# golang
WORKDIR /go-grpc
ENV GO111MODULE on

RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@v1.28
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@v1.2