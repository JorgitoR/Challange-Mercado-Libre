FROM public.ecr.aws/docker/library/golang:1.17-alpine3.14 as builder

ARG GO_MODULES_TOKEN=token

RUN apk update

RUN apk add git
RUN apk add build-base
RUN git config --global url.https://$GO_MODULES_TOKEN@github.com/.insteadOf https://github.com/

ENV GOPRIVATE=github.com/JorgitoR/Challange-Mercado-Libre/*
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -o app cmd/main.go

# Run the Go Binary in Alpine.
FROM public.ecr.aws/docker/library/alpine:3.14

RUN apk update
RUN apk add build-base

WORKDIR /app
COPY --from=builder  /app .
HEALTHCHECK CMD curl --fail http://localhost:8080/healthz || exit 1

EXPOSE 8080 

CMD ["./app"]