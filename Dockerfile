FROM golang:1.18

# Environment variable
ENV ENV_LC_DB_USERNAME="postgres"
ENV ENV_LC_DB_PASSWORD=""
ENV ENV_LC_DB_HOST="127.0.0.1"
ENV ENV_LC_DB_PORT="5432"
ENV ENV_LC_DB_NAME="postgres"
ENV ENV_LC_LISTENING_PORT="6432"

WORKDIR /app/src

# Copy all to workdir
COPY . .

# pre-copy/cache go.mod for pre-downloading dependencies and only redownloading them in subsequent builds if they change
# RUN go mod download && go mod verify
RUN go mod tidy

# Build project
RUN go build -v -o /app ./...

# Change workdir
WORKDIR /app

# Launch app
CMD ["/bin/sh", "-c", "./leasing-car"]