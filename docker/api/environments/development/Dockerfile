FROM golang:1.18

ENV APP_ROOT=/app
ENV GO111MODULE=on

RUN apt update

WORKDIR ${APP_ROOT}

COPY ./server/go.mod ${APP_ROOT}
COPY ./server/go.sum ${APP_ROOT}
RUN go mod tidy

COPY ./server ${APP_ROOT}

RUN go install github.com/cosmtrek/air@latest
RUN go install honnef.co/go/tools/cmd/staticcheck@latest

COPY ./server/start.sh ${APP_ROOT}/start.sh
RUN chmod +x ${APP_ROOT}/start.sh

CMD ["air","-c", ".air.toml"]