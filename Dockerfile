FROM  golang:latest as build-env
ENV APP_NAME sample-dockerize-app
ENV CMD_PATH main.go
ENV GO111MODULE=on
ENV GOOS linux
COPY . $GOPATH/src/$APP_NAME
WORKDIR $GOPATH/src/$APP_NAME
RUN CGO_ENABLED=0 go build -v -o /$APP_NAME $GOPATH/src/$APP_NAME/$CMD_PATH
FROM alpine:3.14
ENV APP_NAME sample-dockerize-app

# Copy only required data into this image
COPY --from=build-env /$APP_NAME .

# Expose application port
EXPOSE 8081

# Start app 
CMD ./$APP_NAME