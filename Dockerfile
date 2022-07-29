FROM golang:latest AS build-stage
WORKDIR /go/src
COPY . /go/src/
RUN go mod download
RUN GOOS=linux GOARCH=amd64 CGO_ENABLED=1 go build -ldflags="-linkmode external -extldflags '-static'" -o gm-exam .

FROM golang:alpine
WORKDIR /go
COPY ./storage/keys/. ./storage/keys/
COPY --from=build-stage /go/src/gm-exam .
CMD ["./gm-exam"] 