#!/bin/bash

# Start Redis Sentinel cluster
echo "Starting Redis Sentinel production setup..."

# Create required directories
mkdir -p config data

# Set proper permissions
chmod 644 config/*.conf

# Start services
docker-compose up -d

# Wait for services to be ready
echo "Waiting for services to be ready..."
sleep 5

# Check status
echo "Checking cluster status..."
docker-compose ps

echo "Redis Sentinel cluster started successfully!"
echo ""
echo "Redis Master: localhost:6379"
echo "Redis Replica 1: localhost:6380"
echo "Redis Replica 2: localhost:6381"
echo "Sentinel 1: localhost:26379"
echo "Sentinel 2: localhost:26380"
echo "Sentinel 3: localhost:26381"