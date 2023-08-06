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
#COPY --from=ghcr.io/tarampampam/curl:8.0.1 /bin/curl /bin/curl
EXPOSE 8000
EXPOSE 80
COPY --from=builder /go/bin/api /go/bin/api
CMD ["/go/bin/api"]
# docker build -t ariefhidayat/gin-gorm-api:0.0.3 .