FROM debian:buster
FROM golang:1.15.7-buster
ENV GO111MODULE=on
ENV GOFLAGS=-mod=vendor
RUN mkdir /app
ADD . /app
ENV APP_HOME /app
WORKDIR ${APP_HOME}
RUN go mod vendor
RUN go mod verify
RUN go build -o pharmApp
EXPOSE 8010
CMD ["./pharmApp"]