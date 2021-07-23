FROM golang:1.15 as builder

COPY . /code
WORKDIR /code
RUN go mod init github.com/facebookincubator/dhcplb && \
    go mod tidy && \
    go build

FROM golang:1.15
EXPOSE 67
COPY --from=builder /code/dhcplb /dhcplb

ENTRYPOINT ["/dhcplb"]
