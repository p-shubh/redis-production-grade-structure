#!/bin/bash

echo "=== Redis Sentinel Cluster Health Check ==="

# Check master status
echo -e "\n1. Checking Master Status:"
docker exec redis-master redis-cli -a redis_password_123 info replication

# Check replica status
echo -e "\n2. Checking Replica 1 Status:"
docker exec redis-replica-1 redis-cli -a redis_password_123 info replication

# Check Sentinel status
echo -e "\n3. Checking Sentinel Status:"
docker exec sentinel-1 redis-cli -p 26379 sentinel masters

# Check Sentinel monitoring
echo -e "\n4. Sentinel Monitoring Details:"
docker exec sentinel-1 redis-cli -p 26379 sentinel slaves mymaster

# Simulate master failure (uncomment to test)
# echo -e "\n5. Testing failover - stopping master..."
# docker stop redis-master
# sleep 15
# docker exec sentinel-1 redis-cli -p 26379 sentinel masters
# docker start redis-master