FROM --platform=${BUILDPLATFORM} golang:1.17-alpine3.15 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app

ARG GOPROXY
# download deps before gobuild
RUN go mod download -x
ARG TARGETOS
ARG TARGETARCH
RUN scripts/container_binary.sh

FROM alpine:3.15

LABEL maintainer="Jeeva Kandasamy <jkandasa@gmail.com>"

ENV APP_HOME="/app"

# install timzone utils
RUN apk --no-cache add tzdata

# create application home
RUN mkdir -p ${APP_HOME}

# copy UI
COPY web-console/build /ui

# copy application bin file
COPY --from=builder /app/youtube-dl ${APP_HOME}/youtube-dl

RUN chmod +x ${APP_HOME}/youtube-dl

WORKDIR ${APP_HOME}

ENTRYPOINT [ "/app/youtube-dl" ]