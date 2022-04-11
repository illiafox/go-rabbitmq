# build stage
FROM golang:1.18-alpine AS build-env
RUN apk --no-cache add build-base git
ADD . /publisher
RUN cd /publisher/cmd/app && go build -o bin

# final stage
FROM alpine
WORKDIR /app
COPY --from=build-env /publisher/cmd/app/bin /app/
ENTRYPOINT ./bin