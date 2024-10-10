FROM golang:1.23-alpine AS builder

WORKDIR /app
COPY . .

RUN go mod download

RUN mkdir -p /usr/bin && CGO_ENABLED=0 go build -o /usr/bin/app ./cmd

FROM alpine

RUN apk add --no-cache tzdata

ENV TZ=Asia/Jakarta
RUN cp /usr/share/zoneinfo/$TZ /etc/localtime && echo $TZ > /etc/timezone

COPY --from=builder /usr/bin/app /usr/bin/app

CMD ["app"]