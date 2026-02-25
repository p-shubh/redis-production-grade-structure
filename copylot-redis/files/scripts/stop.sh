#!/bin/bash

# Navigate to parent directory
cd "$(dirname "$0")/.."

echo "Stopping Redis Sentinel cluster..."
docker-compose down

echo "To remove volumes (WARNING: deletes data):"
echo "docker-compose down -v"

echo "Stopped successfully!"