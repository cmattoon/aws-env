FROM golang:alpine3.7
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main.app .

FROM alpine:latest
RUN mkdir /app
WORKDIR /app
COPY --from=0 /app/main.app /app/main.app
CMD ["/app/main.app"]


