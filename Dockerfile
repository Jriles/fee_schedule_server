FROM golang:1.10 AS build
WORKDIR /go/src
COPY go ./go
COPY main.go .

ENV CGO_ENABLED=0
RUN go get -d -v ./...

RUN go build -a -installsuffix cgo -o fee_schedule_server .

FROM scratch AS runtime
ENV GIN_MODE=release
COPY --from=build /go/src/fee_schedule_server ./
EXPOSE 8080/tcp
ENTRYPOINT ["./fee_schedule_server"]
