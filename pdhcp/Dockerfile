FROM golang:1.15 as builder

COPY . /code
WORKDIR /code
RUN make

EXPOSE 67

FROM alpine

COPY --from=builder /code/pdhcp /pdhcp

ENTRYPOINT ["/pdhcp"]
