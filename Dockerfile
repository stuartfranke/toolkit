FROM golang:1.19-bullseye

ENV CGO_ENABLED 0
#ENV GOPATH /go
#ENV PATH "${PATH}":/go/bin:"${GOPATH}"/bin

COPY ./ /tmp

WORKDIR /tmp
