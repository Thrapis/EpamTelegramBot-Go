FROM mysql:8.0.23

COPY ./migrations/*.sql /docker-entrypoint-initdb.d/