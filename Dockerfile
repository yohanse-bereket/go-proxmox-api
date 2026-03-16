# Use Go official image
FROM golang:1.25-alpine

# Install necessary tools
RUN apk add --no-cache bash git curl unzip

# Install Terraform
ENV TERRAFORM_VERSION=1.14.7
RUN curl -fsSL https://releases.hashicorp.com/terraform/${TERRAFORM_VERSION}/terraform_${TERRAFORM_VERSION}_linux_amd64.zip -o terraform.zip \
    && unzip terraform.zip \
    && mv terraform /usr/local/bin/ \
    && rm terraform.zip

# Set working directory
WORKDIR /app

# Copy go.mod and download dependencies
COPY go.mod go.sum ./
RUN go mod download

# Copy all project files
COPY . .

# Build the Go binary
RUN go build -o server ./main.go

# Expose API port
EXPOSE 8080

# Start the server
CMD ["./server"]
