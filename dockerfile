# FROM golang:latest as builder
# WORKDIR /app
# COPY go.mod go.sum ./
# RUN go mod download
# COPY . .
# RUN go build -o main .
# # EXPOSE 80
# # ENTRYPOINT [ "./main" ]

# FROM gcr.io/distroless/static-debian11
# # RUN mkdir /main
# COPY --from=builder /app/main .
# EXPOSE 80
# CMD ["/main"]

FROM golang:latest as builder

# Set the working directory inside the container
WORKDIR /app

# Copy the go module files and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy the rest of the application source code
COPY . .

# Build the application binary
RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o /app/main .

# Verify the binary exists and has the correct permissions
RUN ls -l /app/main

FROM alpine:latest

# Install CA certificates (needed for HTTPS connections)
RUN apk --no-cache add ca-certificates

# Copy the binary from the builder stage
COPY --from=builder /app/main /main

# Ensure the binary has execute permissions
RUN chmod +x /main

# Expose the port the application will run on
EXPOSE 80

# Set the entry point to the binary
ENTRYPOINT ["/main"]