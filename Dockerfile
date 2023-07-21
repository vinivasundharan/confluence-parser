# Use the official Go image as the base image
FROM golang:latest
WORKDIR /markdown
COPY . .
RUN go mod download && go build
EXPOSE 8080

# Command to run the Go application when the container starts
CMD ["./markdown"]  