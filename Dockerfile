FROM golang AS builder

RUN apt-get update -y \
    && apt-get install -y upx

RUN mkdir /app
COPY /src /app

WORKDIR /app
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -ldflags "-s -w" -o main app/main.go \
    && upx --brute main \
    && ls -l

FROM scratch
COPY --from=builder /app/main /
CMD ["/main"]
