FROM golang:1.18 AS builder
WORKDIR /app
COPY go.mod .
COPY go.sum .
RUN go env -w GOPROXY=direct && go mod download
COPY cmd/ ./cmd/
COPY adapter/ ./adapter/
COPY pkg/ ./pkg/
COPY config/ ./config/
COPY internal/ ./internal/
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/api  ./cmd/server/main.go

FROM scratch
COPY --from=builder /go/bin/api /go/bin/api
EXPOSE 8000
ENTRYPOINT ["/go/bin/api"]
# docker build -t ariefhidayat/gin-gorm-api:0.0.1 .