# ubuntu part
docker_clean:
#	remove "Exited" or "Created" container
	sudo docker rm `sudo docker ps -a | grep -e "Exited" -e "Created" | sed 's/^\([a-z0-9]*\)\s*.*/\1/'`
#	remove <none> images
	sudo docker rmi -f `sudo docker images | egrep "<none>\s+<none>\s+([a-z0-9]+)" | sed 's/^\(<none>\)\s*\(<none>\)\s*\([a-z0-9]*\)\s*.*/\3/'`
#	remove leasing-car images
	sudo docker rmi -f `sudo docker images | egrep "leasing-car\s+[a-z0-9A-Z-]+\s+([a-z0-9]+)" | sed 's/^\(leasing-car\)\s*\([a-z0-9A-Z-]*\)\s*\([a-z0-9]*\)\s*.*/\3/'`

docker_build:
	sudo docker build --tag leasing-car .

docker_run:
	sudo docker run --name leasing-car-app --env-file ./docker.env -e ENV_LC_DB_PASSWORD=${ENV_LC_DB_PASSWORD}  -p 6432:6432 -d leasing-car

docker_stop:
	sudo docker stop leasing-car-app

docker_debug:
	sudo docker exec -it `sudo docker ps | grep "leasing-car" | sed 's/^\([a-z0-9]*\)\s*.*/\1/'` /bin/sh


# BINARY_PATH="./build"
# BINARY_NAME=""
# PLATFORM=windows # windows / darwin / linux

# build:
# 	go build -o ${BINARY_PATH}/ ./...
# #	GOARCH=amd64 GOOS=${PLATFORM} go build -o ${BINARY_PATH}/ ./...

# # ifeq ($(BINARY_NAME),"")
# # 	GOARCH=amd64 GOOS=${PLATFORM} go build -o ${BINARY_PATH}/ ./...
# # else
# # 	GOARCH=amd64 GOOS=${PLATFORM} go build -o ${BINARY_PATH}/${BINARY_NAME} cmd/leasing-car/main.go
# # endif

# run:
# 	${BINARY_PATH}/${BINARY_NAME}

# build_and_run: build run

# clean:
# 	go clean
# #	rm ${BINARY_PATH}/${BINARY_NAME}-windows

# test:
# # 	go test ./... | grep -e "ok" -e "failed"
# 	go test ./...
