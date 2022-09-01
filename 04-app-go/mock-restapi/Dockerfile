# Start from golang alpine base image
FROM golang:alpine as builder

# Set the current working directory
WORKDIR /build

# Populate modules cache based on go.{mod,sum} files
COPY ./go.mod .
COPY ./go.sum .

# Download all dependencies, will be catch if go.mod and go.sum file not changed
RUN go mod download

# Copy everything from current direcotyr to PWD
COPY ./src/ .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o app .


# Start new stage from scatch for a smaller image
FROM scratch

#RUN apk --no-cache add ca-certificates

#WORKDIR /root/

# Copy Pre-built binary file from previous stage
COPY --from=builder /build/app /app

# Expose port 9090 to outside
EXPOSE 9090

# Run the binary application
CMD ["/app"]


