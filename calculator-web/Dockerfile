FROM golang:1.21-alpine AS build

WORKDIR /app

COPY . .

# Build without modules
RUN GO111MODULE=off CGO_ENABLED=0 go build -o calculator .

FROM alpine:latest
WORKDIR /app
COPY --from=build /app/calculator .
COPY --from=build /app/static ./static

EXPOSE 8080
ENTRYPOINT ["./calculator"]

