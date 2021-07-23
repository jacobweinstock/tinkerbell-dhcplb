FROM golang:1.15

COPY . /code
WORKDIR /code
RUN make

EXPOSE 67

ENTRYPOINT ["./pdhcp"]
