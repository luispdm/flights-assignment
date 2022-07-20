FROM golang:1.18.3-buster AS builder

WORKDIR /opt

COPY cmd/http cmd/http
COPY internal internal
COPY go.mod go.mod
COPY go.sum go.sum
COPY Makefile Makefile

RUN make binary

FROM alpine:3.16

WORKDIR /opt

RUN addgroup -g 1000 srvgroup && \
    adduser -D srvuser -u 1000 -G srvgroup

USER srvuser:srvgroup

COPY --from=builder /opt/server server

CMD [ "./server" ]
