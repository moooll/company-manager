FROM golang:1-alpine AS build-env

COPY  . /app

WORKDIR /app
RUN go build

FROM alpine

ENV DB_HOST=
ENV DB_PORT=5432

COPY --from=build-env /app/repo-service ./
COPY --from=build-env /app/docker/service/wait-for ./

EXPOSE 3000

CMD ["sh", "-c", "./wait-for ${DB_HOST}:${DB_PORT} -t 15 -- ./repo-service"]
