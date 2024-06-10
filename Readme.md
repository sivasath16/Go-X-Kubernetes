## Description
Microservices with GoLang with kubernetes deployment


```
docker build -t go-app-n:latest .
docker run -d -p 80:80 --name web go-app-n:latest

docker-compose build
docker-compose up -d
docker-compose ps

kubectl -f apply kubernetes/
```
