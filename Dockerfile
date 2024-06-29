# Use a specific version of Go that matches your requirements
FROM golang:1.22-alpine

# Install curl, git, and PostgreSQL client tools
RUN apk add --no-cache curl git postgresql-client

# Set Go environment variable
ENV PATH /usr/local/go/bin:$PATH

WORKDIR /order_app

# Copy go mod files and install dependencies
COPY go.mod go.sum ./

RUN go mod tidy

# Install sql-migrate
RUN go install github.com/rubenv/sql-migrate/...@latest

# Copy the rest of the application files
COPY . .

# Ensure the migration script is executable
RUN chmod +x ./run_migrations.sh

# Set environment variables
ENV GO_ENV=development

# Expose the application port
EXPOSE 8080

# Run the migration script
CMD ["./run_migrations.sh"]

RUN chmod +x ./run_go.sh
CMD ["./run_go.sh"]


