FROM golang:alpine AS builder

LABEL stage=gobuilder

ARG GITHUB_TOKEN
ARG BUILD_MODE

ENV CGO_ENABLED=0
ENV GOPRIVATE=github.com/exceptionteapots
RUN apk update --no-cache && apk add --no-cache git
RUN go install github.com/swaggo/swag/cmd/swag@latest
# RUN git config --global url."https://exceptionteapots:$GITHUB_TOKEN@github.com/".insteadOf "https://github.com/"

WORKDIR /build

ADD go.mod .
ADD go.sum .
RUN go mod download
COPY . .
RUN if [[ "$BUILD_MODE" == "debug" ]] ; then swag i -g main_debug.go ; fi
RUN go build --tags $BUILD_MODE -ldflags="-s -w" -buildvcs=false -o /app/gRPCexample


FROM scratch


WORKDIR /app
COPY --from=builder /app/gRPCexample /app/gRPCexample

CMD ["./gRPCexample"]
