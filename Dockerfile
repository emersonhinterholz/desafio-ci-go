FROM golang AS builder

RUN apt-get update -y \
    && apt-get install -y upx

RUN mkdir /app
COPY ./src/main.go /app/

WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-s -w" -o main main.go \
    && upx --brute server \
    && ls -l

FROM scratch
COPY --from=builder /app/server /
CMD ["/main"]
