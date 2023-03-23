FROM alpine:latest

COPY bin/gofm /gofm

VOLUMN /data

CMD ["/gofm" "--config=/data/config.yaml"]