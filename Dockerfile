# build stage
FROM golang:1.8-alpine AS build-env
ADD . /go/src/github.com/IssueSquare/money-mom
RUN cd /go/src/github.com/IssueSquare/money-mom && CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app

# final stage
FROM centurylink/ca-certs
COPY --from=build-env /app /

EXPOSE 8080

ENTRYPOINT ["/app"]
