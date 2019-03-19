ARG GO_VERSION=1.12

FROM golang:${GO_VERSION}-alpine AS building

WORKDIR /app

COPY main.go .

RUN CGO_ENABLED=0 GOARCH=amd64 GOOS=linux go build -o /main .


FROM scratch

WORKDIR /app
COPY --from=building /main ./main
COPY ./public /app/public

EXPOSE 8001

CMD ["./main"]