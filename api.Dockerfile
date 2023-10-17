FROM golang:1.19.5-bullseye AS builder

COPY . /dineflow-review-services
WORKDIR /dineflow-review-services
RUN go mod tidy
RUN go build

FROM golang:1.19.5-bullseye AS runner
# ENV GIN_MODE release

RUN mkdir /app
WORKDIR /app
COPY --from=builder /dineflow-review-services/dineflow-review-services /app

EXPOSE 8091

CMD ["/app/dineflow-review-services"]