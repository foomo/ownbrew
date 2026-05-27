FROM alpine:latest

RUN apk --no-cache add ca-certificates

RUN adduser -D -u 1001 -g 1001 ownbrew

COPY ownbrew /usr/bin/

USER ownbrew
WORKDIR /home/ownbrew

ENTRYPOINT ["ownbrew"]
