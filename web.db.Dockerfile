FROM mysql:8.0.23

COPY ./web/db/*.sql /docker-entrypoint-initdb.d/