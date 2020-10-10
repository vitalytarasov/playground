FROM alpine:3.12.0
RUN apk update && apk add --no-cache bash ca-certificates
WORKDIR /app
COPY play .
USER nobody
ENTRYPOINT /app/play
