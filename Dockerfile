FROM golang:1.12.4 AS builder


# Build Args
ARG gitServer=github.com
ARG gitOrg=Benbentwo
ARG gitRepo=k-ates

ENV GIT_SERVER=$gitServer
ENV GIT_ORG=$gitOrg
ENV GIT_REPO=$gitRepo

# Build arguments
ARG binary_name=k-ates
    # See ./sample-data/go-os-arch.csv for a table of OS & Architecture for your base image
ARG target_os=linux
ARG target_arch=amd64

# Build the server Binary
WORKDIR /go/src/${GIT_SERVER}/${GIT_ORG}/${GIT_REPO}
ADD . ./
RUN go get -u ./...
RUN CGO_ENABLED=0 GOOS=${target_os} GOARCH=${target_arch} go build -a -o /app/${binary_name} main.go

RUN ls /app
#FROM scratch

# Default to localhost
# Should match the default set in values.yaml
ARG port=80
ARG root=/
ENV PORT=$port
ENV ROOT=$root

WORKDIR /app/
#COPY --from=builder /app/${binary_name} .
RUN ["chmod", "-R", "+x", "/app"]
EXPOSE ${PORT}

RUN ls -la /app
CMD /app/k-ates
#ENTRYPOINT ["cat", "/dev/null"]
#ENTRYPOINT /app/${binary_name}
#CMD 'go run main.go'
