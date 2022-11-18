# --- base -------------------------------------------------------------------
FROM golang:1.17-alpine as base

RUN apk update && apk add bash

WORKDIR /snapshot
ENV GO111MODULE=on
ENV GOPATH="/go"
ENV PATH="$GOPATH/bin:$PATH"

RUN apk --update add --no-cache git ca-certificates shadow \
    && update-ca-certificates \
    && mkdir -p /home/dockeruser/snapshot/bin \
    && groupadd -r dockeruser \
    && useradd -r -g dockeruser dockeruser

COPY go.mod .
COPY go.sum .
RUN go mod download

# --- builder -------------------------------------------------------------------
FROM base as builder

WORKDIR /snapshot
ARG SERVICE_NAME
ENV SERVICE_NAME=$SERVICE_NAME
ENV CGO_ENABLED=0
ENV GOOS=linux
ENV GOARCH=amd64

COPY api/ api/

RUN go build -o /build /snapshot/api/cmd

# --- build -------------------------------------------------------------
FROM builder as build

RUN go build -tags netgo -ldflags "-X main.version=0.0.0 -w -s -X main.envType=production" -o bin/build /snapshot/api/cmd

# --- deployment -------------------------------------------------------------
FROM scratch as production

COPY --from=build /snapshot/bin/build /
EXPOSE 8000
ENTRYPOINT  ["/build"]

# ---- local ----------------------------------------------------------------------------
FROM builder as local

EXPOSE 8000
ENTRYPOINT  ["/build"]