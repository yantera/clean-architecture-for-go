FROM golang:1.18

ENV APP_ROOT=/app
ENV GO111MODULE=on

RUN apt update

WORKDIR ${APP_ROOT}

COPY ./server/go.mod ${APP_ROOT}
COPY ./server/go.sum ${APP_ROOT}
RUN go mod tidy

COPY ./server ${APP_ROOT}

CMD ["go","run", "main.go"]