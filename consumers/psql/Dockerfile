# build stage
FROM golang:1.18-alpine AS build-env
RUN apk --no-cache add build-base git
ADD . /consumer-psql
RUN cd /consumer-psql/cmd/app && go build -o bin

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /consumer-psql/cmd/app/bin /app/
ENTRYPOINT ./bin