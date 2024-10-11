FROM golang:alpine AS builder

LABEL stage=gobuilder

ARG GITHUB_TOKEN
ARG BUILD_MODE

ENV CGO_ENABLED=0
ENV GOPRIVATE=github.com/exceptionteapots
RUN apk update --no-cache && apk add --no-cache git protoc
RUN go install google.golang.org/protobuf/cmd/protoc-gen-go@latest
RUN go install google.golang.org/grpc/cmd/protoc-gen-go-grpc@latest

# RUN git config --global url."https://exceptionteapots:$GITHUB_TOKEN@github.com/".insteadOf "https://github.com/"

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/helloworld_service.proto
RUN go build --tags $BUILD_MODE -ldflags="-s -w" -buildvcs=false -o /app/gRPCexample


FROM scratch


WORKDIR /app
COPY --from=builder /app/gRPCexample /app/gRPCexample

CMD ["./gRPCexample"]
