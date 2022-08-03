# leasing-car project exercice



<p align="center">
  <img src="https://user-images.githubusercontent.com/22281217/182313702-d44449a0-b7d9-4cf3-bbf2-e778f4fa29af.png">
</p>


## Demo
https://api.leasing-car.r-wg.it/

## Utilisation

### Setting environment variables
#### Windows
```
$env:ENV_LC_DB_USERNAME="postgres"
$env:ENV_LC_DB_PASSWORD="postgres_password"
$env:ENV_LC_DB_HOST="127.0.0.1"
$env:ENV_LC_DB_PORT="5432"
$env:ENV_LC_DB_NAME="postgres"
$env:ENV_LC_LISTENING_PORT="6432"
```
#### Linux
```
export ENV_LC_DB_USERNAME="postgres"
export ENV_LC_DB_PASSWORD="postgres_password"
export ENV_LC_DB_HOST="127.0.0.1"
export ENV_LC_DB_PORT="5432"
export ENV_LC_DB_NAME="postgres"
export ENV_LC_LISTENING_PORT="6432"
```

### PostgreSQL installation
```
# run postgresql:12-alpine image in detach mode
sudo docker run --name postgres-db -e POSTGRES_PASSWORD=$ENV_LC_DB_PASSWORD -p 127.0.0.1:$ENV_LC_DB_PORT:$ENV_LC_DB_PORT -d postgres:12-alpine

# start postgres-db container if it is not
sudo docker start postgres-db

# import init data in postgres-db container
sudo docker exec -i postgres-db psql -U $ENV_LC_DB_USERNAME -d $ENV_LC_DB_NAME < sql/init/init_db.sql

# get IP from the postgres-db container ; for example 172.17.0.2
sudo docker inspect `sudo docker ps | grep "postgres" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'` | grep IPAddress

# update environment variable with ip of db container
export ENV_LC_DB_HOST="172.17.0.2"
```


### Application installation
```
# build leasing-car image
sudo docker build --tag leasing-car .

# run leasing-car-app image in detach mode
sudo docker run \
--name leasing-car-app \
--env-file ./docker.env \
-e ENV_LC_DB_HOST=$ENV_LC_DB_HOST \
-e ENV_LC_DB_PASSWORD=$ENV_LC_DB_PASSWORD \
-p 127.0.0.1:$ENV_LC_LISTENING_PORT:$ENV_LC_LISTENING_PORT \
-d \
leasing-car

# start postgres-db container if it is not
sudo docker start leasing-car-app

# application test
curl --request GET "http://127.0.0.1:${ENV_LC_LISTENING_PORT}/customer/getall"
```



## TODO

* struct or suppl. var to feeback status and error from sql to route (status for HTTP ; error for log)
* check customer with regexp
* update Makefile
* write requests in a swagger API file
* add postman tests



## Notes

### Docker: installation instruction - specific to Windows (WSL)
```
sudo apt-get update
sudo apt-get install     ca-certificates     curl     gnupg     lsb-release
sudo mkdir -p /etc/apt/keyrings
curl -fsSL https://download.docker.com/linux/ubuntu/gpg | sudo gpg --dearmor -o /etc/apt/keyrings/docker.gpg
echo   "deb [arch=$(dpkg --print-architecture) signed-by=/etc/apt/keyrings/docker.gpg] https://download.docker.com/linux/ubuntu \
$(lsb_release -cs) stable" | sudo tee /etc/apt/sources.list.d/docker.list > /dev/null
sudo apt-get update
sudo apt-get install docker-ce docker-ce-cli containerd.io docker-compose-plugin

sudo service docker start
```

### Docker: debug
```
sudo docker exec -it `sudo docker ps | grep "postgres-db" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'` /bin/sh

sudo docker exec -it `sudo docker ps | grep "leasing-car" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'` /bin/sh

sudo docker run -it --env-file ./docker.env -e ENV_LC_DB_HOST=$ENV_LC_DB_HOST -e ENV_LC_DB_PASSWORD=$ENV_LC_DB_PASSWORD  --entrypoint /bin/sh leasing-car

# get IP from the container
sudo docker inspect `sudo docker ps | grep "postgres" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'` | grep IPAddress
```
