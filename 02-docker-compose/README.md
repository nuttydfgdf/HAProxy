
## Exploring the HAProxy Stats Page


## Start Docker-compose
```
docker-compose up -d
docker-compose ps
docker-compose logs
```

# HAProxy Monitoring 
http://localhost:1936/


# Simple loadtest
ab -c 10 -n 100 http://localhost:8080/

-n: Number of requests
-c: Number of concurrent requests