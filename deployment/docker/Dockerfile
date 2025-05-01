FROM golang:1.24.2-alpine3.20

WORKDIR /app
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -o /server
CMD /server
