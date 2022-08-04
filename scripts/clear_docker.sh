#!/bin/sh

PERSONAL_PATH_APP="/personal/path"

RED='\033[0;31m'
WHITE='\033[0m'

print_red()
{
    echo "$(date) - ${RED}$1${WHITE}"
}

cd $PERSONAL_PATH_APP

print_red "deleting directory leasing-car..."
rm -Rf ./leasing-car

print_red "stopping docker container leasing-car..."
sudo docker stop `sudo docker ps -a | grep "leasing-car" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'`
print_red "deleting docker container leasing-car..."
sudo docker rm `sudo docker ps -a | grep "leasing-car" | sed 's/^\([a-z0-9]*\)\s*.*$/\1/'`
print_red "deleting docker image leasing-car..."
sudo docker rmi -f `sudo docker images | egrep "leasing-car\s+[a-z0-9A-Z-]+\s+([a-z0-9]+)" | sed 's/^\(leasing-car\)\s*\([a-z0-9A-Z-]*\)\s*\([a-z0-9]*\)\s*.*/\3/'`
