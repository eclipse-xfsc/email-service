FROM golang:1.23-alpine3.21

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server
CMD /server
