FROM golang:onbuild
RUN mkdir /app
ADD . /app/
WORKDIR /app
RUN go build -o main.app .

FROM golang:latest
COPY --from=0 /app/main.app /app/main.app
CMD ["/app/main.app"]


