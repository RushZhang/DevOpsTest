#! /bin/sh

kill -9 $(pgrep webserver)
cd ~/DevOpsTest/
git pull https://github.com/RushZhang/DevOpsTest.git
cd webserver
./webserver &
#&是为了让程序在后台运行