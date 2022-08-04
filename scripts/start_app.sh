#!/bin/bash

PERSONAL_PATH_APP="/personal/path"

print_date()
{
    echo "$(date) - $1" >> $PERSONAL_PATH_APP/leasing-app.log
}

DOCKER_ID_DB="7638c95dfa19"
DOCKER_ID_APP="07da3122a222"

print_date "Launching docker containers..."

# start postgres
sudo docker start $DOCKER_ID_DB

print_date "Sleeping a while (30sec)..."
sleep 30

# start leasing-car app
sudo docker start $DOCKER_ID_APP

print_date "Ending launch"
