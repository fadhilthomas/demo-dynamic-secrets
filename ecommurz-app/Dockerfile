FROM golang:1.19-alpine AS build
RUN apk --no-cache add build-base
ADD . /src
RUN cd /src && GOOS=linux GOARCH=arm64 go build -o main main.go

# final stage
FROM alpine:latest
WORKDIR /app
COPY --from=build /src/main /app/
RUN chmod +x /app/main
CMD ["./main"]
EXPOSE 8090
