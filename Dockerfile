##
## Build
##
FROM golang:1.17 AS build
WORKDIR /app
COPY . ./
RUN go mod download
COPY *.go ./
RUN go build -o /project-http-basic

##
## Deploy
##
FROM nginx
WORKDIR /app
COPY --from=build /project-http-basic /project-http-basic
EXPOSE 8000
ENTRYPOINT ["/project-http-basic"]