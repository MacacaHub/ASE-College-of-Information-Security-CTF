#!/bin/sh

PWN_CHALLENGE_PATH=$1
PWN_CHALLENGE_NAME=$(basename $PWN_CHALLENGE_PATH)

echo "PWN_CHALLENGE_PATH=$PWN_CHALLENGE_PATH"
echo "PWN_CHALLENGE_NAME=$PWN_CHALLENGE_NAME\n"

cp -r $PWN_CHALLENGE_PATH/src src

docker build -t pwn-builder . >/dev/null
docker run -d  -v $(pwd)/src:/src --name pwn-builder pwn-builder make >/dev/null

echo "Build logs"
docker logs pwn-builder
docker rm -f pwn-builder >/dev/null

chmod +x src/$PWN_CHALLENGE_NAME
cp src/$PWN_CHALLENGE_NAME $PWN_CHALLENGE_PATH/service/public/
rm -r src
