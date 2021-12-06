# syntax=docker/dockerfile:experimental

FROM golang:1.17

ENV APP_ROOT /app
WORKDIR ${APP_ROOT}

COPY . .

COPY go.mod .
COPY go.sum .
RUN ls
RUN go mod download
