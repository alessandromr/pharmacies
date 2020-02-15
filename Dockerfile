FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git
WORKDIR /go/src/app
COPY . .
RUN go mod download
WORKDIR /go/src/app/cmd/jsonrpc/
RUN GOOS=linux go build -ldflags="-s -w" -o ./main .

FROM alpine:3.11
RUN apk --no-cache add ca-certificates
WORKDIR /usr/bin
COPY --from=build /go/src/app/cmd/jsonrpc/main /go/bin/main
EXPOSE 8080
ENTRYPOINT /go/bin/main