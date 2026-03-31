import os

env = os.getenv('APP_ENV', 'not_set')

print(f"----------------------------------------")
print(f"Ứng dụng đang chạy trong môi trường: {env}")
print(f"----------------------------------------")