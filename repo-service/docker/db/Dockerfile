FROM postgres:13-alpine

COPY ./script.sh /docker-entrypoint-initdb.d

EXPOSE 5432
CMD ["postgres"]
