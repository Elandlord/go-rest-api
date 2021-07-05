FROM golang:1.16.5-alpine
RUN mkdir /work
ADD . /work
WORKDIR /work
RUN go mod download
RUN go build -o main .
CMD ["/work/main"]