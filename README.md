# leasing-car project exercice



<p align="center">
  <img src="https://user-images.githubusercontent.com/22281217/182313702-d44449a0-b7d9-4cf3-bbf2-e778f4fa29af.png">
</p>



## Utilisation

### PostgreSQL installation
```
docker run --name postgres-db -e POSTGRES_PASSWORD=$PASSWORD_DOCKER_POSTGRESQL_ROOT -p 5432:5432 -d postgres:12-alpine

docker start postgres-db

docker inspect `docker ps | grep "postgres" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'` | grep IPAddress
```


### Application installation
```
docker build --tag leasing-car .

export ENV_LC_DB_PASSWORD="my_password"

docker run --name leasing-car-app --env-file ./docker.env -p 6432:6432 -d leasing-car
```


## TODO

* struct to feeback status and error from sql to route (status for HTTP ; error for log)
* check customer with regexp
* update Makefile
* update README.md with commands to launch the project
* write requests in a swagger API file
* add postman tests



## Notes

### Docker: installation instruction - specific to Windows (WSL)
```
sudo apt-get install docker
sudo apt-get update && sudo apt-get upgrade -y
sudo apt-get install     ca-certificates     curl     gnupg     lsb-release
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo   "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin

service docker start
```

### Docker: debug
```
docker exec -it `docker ps | grep "leasing-car" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'` /bin/sh
docker run -it --env-file ./docker.env --entrypoint /bin/sh leasing-car
```
