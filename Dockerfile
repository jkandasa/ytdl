FROM --platform=${BUILDPLATFORM} golang:1.23-alpine3.21 AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app

# install timzone utils
RUN apk --no-cache add tzdata

ARG GOPROXY
# download deps before gobuild
RUN go mod download -x
ARG TARGETOS
ARG TARGETARCH
RUN scripts/container_binary.sh

FROM alpine:3.21

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