FROM docker.io/library/golang:alpine@sha256:b35984144ec2c2dfd6200e112a9b8ecec4a8fd9eff0babaff330f1f82f14cb2a 

RUN mkdir /proto

RUN mkdir /stubs

RUN apk -U --no-cache add git protobuf

RUN apk add bash

ENV PATH="/go/src/github.com/tokopedia/gripmock:${PATH}"

RUN go get -u -v github.com/golang/protobuf/protoc-gen-go \
	google.golang.org/grpc \
	google.golang.org/grpc/reflection \
	golang.org/x/net/context \
	github.com/go-chi/chi \
	github.com/lithammer/fuzzysearch/fuzzy \
	golang.org/x/tools/imports

RUN go get github.com/markbates/pkger/cmd/pkger

# cloning well-known-types
RUN git clone --depth=1 https://github.com/google/protobuf.git /protobuf-repo

RUN mkdir protobuf

# only use needed files
RUN mv /protobuf-repo/src/ /protobuf/

RUN rm -rf /protobuf-repo

RUN mkdir -p /go/src/github.com/tokopedia/gripmock

COPY . /go/src/github.com/tokopedia/gripmock

RUN chmod u+x /go/src/github.com/tokopedia/gripmock/fix_gopackage.sh

WORKDIR /go/src/github.com/tokopedia/gripmock/protoc-gen-gripmock

RUN pkger

# install generator plugin
RUN go install -v

WORKDIR /go/src/github.com/tokopedia/gripmock

# install gripmock
RUN go install -v

# remove all .pb.go generated files
# since generating go file is part of the test
RUN find . -name "*.pb.go" -delete -type f

EXPOSE 4770 4771

ENTRYPOINT ["gripmock"]
