# Build the application
FROM golang:1.21

WORKDIR /app

# Copy and download dependencies
# COPY go.mod go.sum ./
# RUN go mod download


COPY *.go ./

# RUN rm go.mod && rm go.sum && rm go.work 

RUN cd /app

RUN go mod init curdusers

RUN go mod tidy

RUN go mod download

# RUN go get "curdusers/configs"

# RUN go get "curdusers/routes"

# # Copy source code
# COPY *.go ./

# # Build the application
# RUN go build -o /docker-gs-ping

# # Create a new stage for the final image
# FROM golang:1.21

# WORKDIR /

# # Copy the binary from the build-stage
# COPY --from=build-stage /docker-gs-ping /docker-gs-ping

EXPOSE 8080

# Specify the user if needed
# USER nonroot:nonroot

# Define the entry point
ENTRYPOINT ["go", "run", "main.go"]
