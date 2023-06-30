FROM golang:latest
COPY . /usr/local/app
WORKDIR /usr/local/app
RUN go mod init local/app
