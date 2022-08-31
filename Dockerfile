FROM public.ecr.aws/docker/library/golang:1.17-alpine3.14 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/server/main.go

FROM alpine:latest AS production
COPY --from=builder /app .
CMD ["./app"]