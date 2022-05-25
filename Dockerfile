FROM golang:1.16.0-stretch

WORKDIR /go/src
ENV PATH="/go/bin:${PATH}"
ENV CGO_ENABLED=0

CMD ["tail", "-f", "/dev/null"]
# syntax=docker/dockerfile:1

# FROM golang:1.16-alpine

# WORKDIR /app

# COPY go.mod ./
# COPY go.sum ./

# RUN go mod download

# COPY *.go ./

# RUN go build -o /docker-gs-ping

# EXPOSE 8181

# CMD [ "/docker-gs-ping" ]
# CMD [ "go", "run", "main.go" ]