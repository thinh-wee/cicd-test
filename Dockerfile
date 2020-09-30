######## Start from the latest golang base image #######
FROM amd64/golang:1.15-alpine

# Add Maintainer Info
LABEL maintainer="WeeDigital Company | admin@wee.vn"

RUN apk --no-cache add tzdata

ARG TZ=Asia/Ho_Chi_Minh
RUN ln -fs /usr/share/zoneinfo/${TZ} /etc/localtime;\
    date

# Set the Current Working Directory inside the container
WORKDIR /app

# Copy go mod and sum files
COPY go.mod go.sum ./

# Download all dependencies. Dependencies will be cached if the go.mod and go.sum files are not changed
RUN go mod download

# Copy the source from the current directory to the Working Directory inside the container
COPY . .

# Build the Go app and clean resources
RUN CGO_ENABLE=0 go build -ldflags "-extldflags \"-static\" -s -w" -o bin/application -trimpath api-gateway/*.go

# This container exposes port 8080 to the outside world
EXPOSE 8080

ENTRYPOINT [ "bin/application" ]