FROM golang:1.20-alpine AS build

WORKDIR /app

COPY go.mod go.sum ./

RUN make deps

COPY . .

RUN make carecompanion

FROM alpine:latest

WORKDIR /app

COPY --from=build /app/carecompanion .

EXPOSE 8888

CMD ["./carecompanion", "serve"]
