# build stage
FROM golang:1.18-alpine AS build-env
RUN apk --no-cache add build-base git
ADD . /consumer-redis
RUN cd /consumer-redis/cmd/app && go build -o bin

# final stage
FROM alpine
WORKDIR /app

EXPOSE $API_PORT

COPY --from=build-env /consumer-redis/cmd/app/bin /app/
ENTRYPOINT ./bin