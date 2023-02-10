FROM golang:1.20-alpine

WORKDIR /usr/src/app

ENTRYPOINT ["docker-entrypoint.sh"]
COPY ./entrypoint.sh /usr/local/bin/docker-entrypoint.sh

RUN chmod +x /usr/local/bin/docker-entrypoint.sh

COPY . .
RUN go build -v -o /usr/local/bin/app .



EXPOSE 9100

STOPSIGNAL SIGINT

CMD ["app"]
