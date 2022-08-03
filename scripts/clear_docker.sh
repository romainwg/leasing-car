#!/bin/sh

# Personal script automatisation

RED='\033[0;31m'
WHITE='\033[0m'
ROOT_PATH_APP="/path/to/my/folder"

cd $ROOT_PATH_APP

echo "${RED}deleting directory leasing-car...${WHITE}"
rm -Rf ./leasing-car

echo "${RED}stopping docker container leasing-car...${WHITE}"
sudo docker stop `sudo docker ps -a | grep "leasing-car" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'`
echo "${RED}deleting docker container leasing-car...${WHITE}"
sudo docker rm `sudo docker ps -a | grep "leasing-car" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'`
echo "${RED}deleting docker image leasing-car...${WHITE}"
sudo docker rmi -f `sudo docker images | egrep "leasing-car\s+[a-z0-9A-Z-]+\s+([a-z0-9]+)" | sed 's/^\(leasing-car\)\s*\([a-z0-9A-Z-]*\)\s*\([a-z0-9]*\)\s*.*/\3/'`
