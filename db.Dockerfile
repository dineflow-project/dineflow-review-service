# start with base image
FROM mysql:latest

COPY ./database/*.sql /docker-entrypoint-initdb.d/