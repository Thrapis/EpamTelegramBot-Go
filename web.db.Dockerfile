FROM mysql:8.0.23

COPY ./sql/*.sql /docker-entrypoint-initdb.d/