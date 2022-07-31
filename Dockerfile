FROM golang:1.18-alpine as BUILD

WORKDIR /app/
COPY . /app/