FROM golang:1.16-alpine3.14

WORKDIR /src/

RUN mkdir api.server-files

WORKDIR /src/api.server-files

ENV KEY=<KEY>
ENV HOSTNAME=<ip>
ENV PORT=<port>

COPY . .

RUN go build -o app src/main.go

EXPOSE 8001

CMD ["./app"]