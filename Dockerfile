FROM golang:alpine AS build
RUN apk --no-cache add gcc g++ make
RUN apk add git ca-certificates && update-ca-certificates
WORKDIR /go/src/app
COPY . .
RUN go mod download
WORKDIR /go/src/app/cmd/jsonrpc/
RUN GOOS=linux CGO_ENABLED=0 go build -ldflags="-s -w" -o ./main .

FROM scratch
COPY --from=build /etc/ssl/certs/ca-certificates.crt /etc/ssl/certs/
COPY --from=build /go/src/app/cmd/jsonrpc/main /go/src/app/cmd/jsonrpc/main
EXPOSE 8080
ENTRYPOINT ["/go/src/app/cmd/jsonrpc/main"]