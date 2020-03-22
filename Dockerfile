FROM golang:latest

ENV GO111MODULE=auto
ENV GOPROXY=https://goproxy.cn,https://goproxy.io,direct

WORKDIR /app
COPY . /app
RUN go build .

EXPOSE 8000
ENTRYPOINT ["./gin-demo"]

