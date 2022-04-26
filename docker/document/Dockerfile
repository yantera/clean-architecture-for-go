FROM golang:1.18

ENV APP_ROOT=/app
ENV GO111MODULE=on

RUN apt update

WORKDIR ${APP_ROOT}

COPY ./documents/go.mod ${APP_ROOT}
COPY ./documents/go.sum ${APP_ROOT}
RUN go mod tidy

COPY ./documents ${APP_ROOT}

RUN go install goa.design/goa/v3/cmd/goa@v3