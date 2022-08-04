#!/bin/sh

PERSONAL_PATH_APP="/personal/path"
START_APP_SCRIPT="start_app.sh"

RED='\033[0;31m'
WHITE='\033[0m'

print_red()
{
    echo "$(date) - ${RED}$1${WHITE}"
}

cd $PERSONAL_PATH_APP

print_red "Cloning repository..."
git clone https://github.com/romainwg/leasing-car

print_red "Changing directory..."
cd leasing-car

print_red "Building image..."
# build leasing-car image
sudo docker build --tag leasing-car .

print_red "Showing environment variable..."
echo $ENV_LC_DB_USERNAME
echo $ENV_LC_DB_PASSWORD
echo $ENV_LC_DB_HOST
echo $ENV_LC_DB_PORT
echo $ENV_LC_DB_NAME
echo $ENV_LC_LISTENING_PORT

print_red "Running image to container..."
# run leasing-car-app image in detach mode
sudo docker run \
--name leasing-car-app \
--env-file ./docker.env \
-e ENV_LC_DB_HOST=$ENV_LC_DB_HOST \
-e ENV_LC_DB_PASSWORD=$ENV_LC_DB_PASSWORD \
-p $ENV_LC_LISTENING_PORT:$ENV_LC_LISTENING_PORT \
-d \
leasing-car

print_red "Sleeping a while..."
sleep 10

print_red "Changing directory..."
cd ../

print_red "Showing docker container ps status"
sudo docker ps -a
# start postgres-db container if it is not
# sudo docker start leasing-car-app

print_red "Getting docker container id of leasing-car..."
APP_ID=`sudo docker ps -a | grep "leasing-car" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'`

print_red "Getting docker container id of postgres-db..."
DB_ID=`sudo docker ps -a | grep "postgres-db" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'`

print_red "Updating launch script with docker container id..."
sed -i "s/DOCKER_ID_DB=\"[a-z0-9]*\"/DOCKER_ID_DB=\"${DB_ID}\"/g" $START_APP_SCRIPT
sed -i "s/DOCKER_ID_APP=\"[a-z0-9]*\"/DOCKER_ID_APP=\"${APP_ID}\"/g" $START_APP_SCRIPT

print_red "Launching app if it is not..."
./$START_APP_SCRIPT

print_red "Testing app with curl..."
# application test
curl --request GET "http://127.0.0.1:${ENV_LC_LISTENING_PORT}/customer/getall"
