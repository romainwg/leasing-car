#!/bin/sh

# Personal script automatisation

RED='\033[0;31m'
WHITE='\033[0m'
ROOT_PATH_APP="/path/to/my/folder"

cd $ROOT_PATH_APP

echo "${RED}Cloning repository...${WHITE}"
git clone https://github.com/romainwg/leasing-car

echo "${RED}Changing directory...${WHITE}"
cd leasing-car

echo "${RED}Building image...${WHITE}"
# build leasing-car image
sudo docker build --tag leasing-car .

echo "${RED}Showing environment variable...${WHITE}"
echo $ENV_LC_DB_USERNAME
echo $ENV_LC_DB_PASSWORD
echo $ENV_LC_DB_HOST
echo $ENV_LC_DB_PORT
echo $ENV_LC_DB_NAME
echo $ENV_LC_LISTENING_PORT

echo "${RED}Running image to container...${WHITE}"
# run leasing-car-app image in detach mode
sudo docker run \
--name leasing-car-app \
--env-file ./docker.env \
-e ENV_LC_DB_HOST=$ENV_LC_DB_HOST \
-e ENV_LC_DB_PASSWORD=$ENV_LC_DB_PASSWORD \
-p $ENV_LC_LISTENING_PORT:$ENV_LC_LISTENING_PORT \
-d \
leasing-car

echo "${RED}Sleeping a while...${WHITE}"
sleep 10

echo "${RED}Showing docker container ps status${WHITE}"
sudo docker ps -a
# start postgres-db container if it is not
# sudo docker start leasing-car-app

echo "${RED}Testing app with curl...${WHITE}"
# application test
curl --request GET "http://127.0.0.1:${ENV_LC_LISTENING_PORT}/customer/getall"
