# Parses a genesis file from the /callisto/genesis.json path during first initialization
# Set /etc/callisto as home

FROM golang:1.21-alpine AS builder

ARG CALLISTO_VERSION="cosmos/v0.50.x"

RUN apk add --no-cache make git &&\
    cd /tmp &&\
    git clone https://github.com/forbole/callisto.git &&\
    cd callisto &&\
    git checkout $CALLISTO_VERSION &&\
    go mod download &&\
    make build &&\
    apk del git make

FROM alpine:latest

COPY callisto-entrypoint.sh /usr/local/bin/entrypoint.sh
COPY --from=builder /tmp/callisto/build/callisto /usr/local/bin/callisto

# -S is a system user
RUN addgroup -S callisto \
    && adduser -S -G callisto callisto \
    && mkdir /etc/callisto \
    && chown callisto:callisto /etc/callisto \
    && mkdir /callisto
USER callisto

# Prometheus sink port
EXPOSE 8000 

# Hasura actions port
EXPOSE 3000 

ENTRYPOINT ["entrypoint.sh"]
CMD ["callist", "start", "--home /etc/callisto"]
