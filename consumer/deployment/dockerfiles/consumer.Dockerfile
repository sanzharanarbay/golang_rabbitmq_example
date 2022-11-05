FROM golang:1.19-alpine as builder

WORKDIR /consumer

# add some necessary packages
# Adding git, bash and openssh to the image
RUN apk update && apk upgrade && \
    apk add --no-cache bash git openssh

# prevent the re-installation of vendors at every change in the source code
COPY go.mod go.sum ./
RUN go mod tidy && go mod download

# Copy and build the app
COPY . .

RUN go mod tidy && go mod vendor

RUN CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o consumer ./cmd/main.go

# Start a new stage from scratch
FROM  alpine:latest
RUN apk --no-cache add ca-certificates

WORKDIR /root/

# Copy the Pre-built binary file from the previous stage. Observe we also copied the .env file
COPY --from=builder /consumer/consumer .
COPY --from=builder /consumer/.env .

#Command to run the executable
CMD ["./consumer"]
