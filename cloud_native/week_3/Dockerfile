FROM golang:1.17-alpine AS build

WORKDIR /go/src/http

COPY . .
RUN CGO_ENABLED=0  GOOS=linux GOARCH=amd64 go build -o main

FROM scratch
COPY --from=build /go/src/http/main /main

EXPOSE 8080

ENTRYPOINT ["/main"]