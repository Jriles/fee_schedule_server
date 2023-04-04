# Use an official Golang runtime as a parent image
FROM golang:1.17-alpine3.14

# Set the working directory to /go/src/app
WORKDIR /go/src/app

# Install Git
RUN apk update && apk add --no-cache git

# Clone the server repository
RUN git clone https://github.com/Jriles/fee_schedule_server.git

# Set the working directory to /go/src/app/fee_schedule_server
WORKDIR /go/src/app/fee_schedule_server

# Copy the server's go.mod files to the container
COPY go.mod .
COPY go.sum .

# Download Go dependencies
RUN go mod download

# Copy the rest of the server code from the server repo to the container
COPY . .

# Build the server
RUN go build

# Set the environment variable for the database connection string
ENV FEE_SCHEDULE_SERVER_DB_CONN="host=localhost user=postgres password=password dbname=fee_schedule sslmode=disable"

EXPOSE 8080

# Start the server and the frontend
CMD ["sh", "-c", "cd /go/src/app/fee_schedule_server && ./fee_schedule_server"]
