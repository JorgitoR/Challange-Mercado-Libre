FROM public.ecr.aws/docker/library/golang:1.17-alpine3.14 as builder

RUN mkdir /app
ADD . /app
WORKDIR /app

RUN GOOS=linux GOARCH=arm64 go build -tags musl -tags dynamic cmd/main.go 

# Run the Go Binary in Alpine.
FROM public.ecr.aws/docker/library/alpine:3.14

FROM alpine:latest AS production
COPY --from=builder /app .
RUN chmod +x ./main
USER appidentity
HEALTHCHECK CMD curl --fail http://localhost:8080/healthz || exit 1

EXPOSE 8080 

CMD ["./app"]
