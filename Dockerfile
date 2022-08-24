# ===== build go binary =====
FROM golang:1.18.3-alpine as go-builder

WORKDIR /go/src/github.com/yagikota/clean_architecture_wtih_go/

COPY cmd/sample_app/main.go .
COPY go.mod .
COPY go.sum .
RUN mkdir pkg
COPY pkg/ ./pkg

RUN go mod download

RUN go build -o server main.go

# ==== build docker image ====
FROM alpine

RUN apk --no-cache add tzdata

COPY --from=go-builder /go/src/github.com/yagikota/clean_architecture_wtih_go//server server

ENTRYPOINT ["/server"]
