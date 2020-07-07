FROM golang:1.13.4 AS builder

ARG APPLICATION=feed
WORKDIR /go/src/${APPLICATION}
COPY . /go/src/${APPLICATION}
RUN export GO111MODULE=on;CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /go/bin/${APPLICATION} -mod=vendor app/main.go

FROM alpine:3.8 AS release
ARG APPLICATION=feed
# RUN apk update \
#        && apk upgrade \
#        && apk add --no-cache \
#        ca-certificates \
#        && update-ca-certificates 2>/dev/null || true
RUN apk add --no-cache tzdata
ENV TZ=Asia/Shanghai

WORKDIR /var/lib/${APPLICATION}
COPY --from=builder /go/bin/${APPLICATION} /bin/${APPLICATION}
COPY --from=builder /go/src/${APPLICATION}/config /var/lib/${APPLICATION}/config/
COPY --from=builder /go/src/${APPLICATION}/tmp /var/lib/${APPLICATION}/tmp

EXPOSE 9000
ENTRYPOINT ["/bin/mpc"]