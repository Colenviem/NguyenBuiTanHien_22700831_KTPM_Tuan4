# 1. Sử dụng redis:latest làm base image
FROM redis:latest

# 2. Tạo thư mục chứa cấu hình bên trong container
WORKDIR /usr/local/etc/redis

# 3. Copy file cấu hình từ máy host vào container
COPY redis.conf .

# 4. EXPOSE cổng mặc định của Redis
EXPOSE 6379

# 5. Chạy Redis server với file cấu hình vừa copy vào
# Lệnh mặc định của Redis image là "redis-server", 
# ta thêm đường dẫn file cấu hình vào sau đó.
CMD ["redis-server", "/usr/local/etc/redis/redis.conf"]