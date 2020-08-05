FROM golang:1.14-alpine

# Install git

RUN apk update && apk add --no-cache git

WORKDIR /app

# Copy go mod and sum files 
# We want to populate the module cache based on the go.{mod,sum} files.
COPY go.mod go.sum ./
COPY api/go.mod ./

# Download all dependencies. Dependencies will be cached if the go.mod and the go.sum files are not changed 
RUN go mod download 

# Copy the source from the current directory to the working Directory inside the container 
COPY . .

# Run tests
CMD CGO_ENABLED=0 go test -v  ./...