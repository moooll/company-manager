FROM golang:1-alpine AS build-env

COPY . /app

WORKDIR /app
RUN go build

FROM alpine

COPY --from=build-env /app/api-service .

EXPOSE 3000
CMD ["./api-service"]
