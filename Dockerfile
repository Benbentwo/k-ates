FROM golang:1.12.4 AS builder

# Default to localhost
ARG port=80
ENV PORT=$port
ARG root=/
ENV ROOT=$root

# Build arguments
ARG binary_name=k-ates
    # See ./sample-data/go-os-arch.csv for a table of OS & Architecture for your base image
ARG target_os=linux
ARG target_arch=amd64

# Build the server Binary
WORKDIR /app/
ADD . ./
RUN CGO_ENABLED=0 GOOS=${target_os} GOARCH=${target_arch} go build -a -o ${binary_name} main.go

FROM scratch

WORKDIR /app/
COPY --from=builder /app/k-ates .
EXPOSE ${PORT}

CMD ls -al
#ENTRYPOINT bash
#CMD 'go run main.go'
