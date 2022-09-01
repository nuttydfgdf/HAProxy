docker-compose build mock-restapi

> Fix, If you found error please run 'go mod tidy'

For remove none image
```
docker rmi $(docker images --filter "dangling=true" -q --no-trunc)
```

docker-compose up -d 
docker-compose ps
docker-compose logs

Try mock API
http://localhost:9090
http://localhost:9090/users

