# syntax=docker/dockerfile:1

FROM golang:1.21

# Set destination for COPY
WORKDIR /app

# Download Go modules
COPY go.mod ./
RUN go mod download

# Copy the source code. Note the slash at the end, as explained in
# https://docs.docker.com/engine/reference/builder/#copy
COPY *.go ./

RUN apt-get update && apt-get install -y --no-install-recommends \
iproute2 \
# ifconfig \
iputils-ping \
&& \
apt-get clean && \
rm -rf /var/lib/apt/lists/*

# Build
RUN CGO_ENABLED=0 GOOS=linux go build -o /docker-server

# Optional:
# To bind to a TCP port, runtime parameters must be supplied to the docker command.
# But we can document in the Dockerfile what ports
# the application is going to listen on by default.
# https://docs.docker.com/engine/reference/builder/#expose
# EXPOSE 8080

# Run
CMD ["/docker-server"]