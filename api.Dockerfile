FROM golang:1.19.5-bullseye AS builder

COPY . /dineflow-review-service
WORKDIR /dineflow-review-service
RUN go mod tidy
RUN go build

FROM debian:bullseye-slim
ENV GIN_MODE release

RUN mkdir /app
WORKDIR /app
COPY --from=builder /dineflow-review-service/dineflow-review-service /app

EXPOSE 8091

CMD ["/app/dineflow-review-service"]