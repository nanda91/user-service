FROM golang:1.17-alpine AS builder

RUN apk update && apk add --no-cache tzdata \
    && cp /usr/share/zoneinfo/Asia/Jakarta /etc/localtime \
    && echo "Asia/Jakarta" >  /etc/timezone

# Set necessary environmet variables needed for our image
ENV GO111MODULE=on \
    CGO_ENABLED=0 \
    GOOS=linux \
    GOARCH=amd64

# Move to working directory /build
WORKDIR /app/user-service

# Copy and download dependency using go mod
COPY go.mod .
COPY go.sum .
RUN go mod download

# Copy the code into the container
COPY . .

# Build the application
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -installsuffix cgo -o user-service

#second stage
FROM alpine:latest
WORKDIR /root/user-service/
RUN apk update


# Import from builder.
COPY --from=builder /app/user-service .
COPY --from=builder /usr/share/zoneinfo /usr/share/zoneinfo
ENV TZ=Asia/Jakarta

# This container exposes port 8080 to the outside world
EXPOSE 8888

CMD ["./user-service"]