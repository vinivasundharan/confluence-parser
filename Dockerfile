# Use the official Go image as the base image
FROM golang:latest
ENV GO111MODULE=on
WORKDIR app
COPY --chown=0:0 . .
WORKDIR markdown
RUN go build -o conf2md ./main.go
RUN chmod +x conf2md
EXPOSE 8080
CMD ["./conf2md"]  