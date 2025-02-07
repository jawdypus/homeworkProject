# Stage 1: Build the frontend
FROM node:22.11.0-alpine AS frontend-builder

WORKDIR /app

COPY frontend/package*.json ./

RUN npm install

COPY frontend/ .

RUN npm run build

# Stage 2: Build the backend
FROM golang:1.23.3-alpine AS backend-builder

RUN apk add --no-cache chromium

WORKDIR /app

COPY backend/go.mod backend/go.sum ./

RUN go mod download

COPY backend/ .

# Copy the built frontend assets from the previous stage
COPY --from=frontend-builder /app/dist ./frontend

RUN go build -o main .

# Stage 3: Create the final image
FROM alpine:latest

RUN apk add --no-cache chromium

WORKDIR /app

# Copy the backend executable and frontend assets
COPY --from=backend-builder /app .

# Set environment variables
ENV STATIC_FILES_PATH=/app/frontend

EXPOSE 8080

# Command to run the backend
CMD ["./main"]