#!/bin/bash

conf=conf/config.json
log=logs/gin-msgboard-out.log
_pid=`ps -ef | grep -v grep|grep gin-msgboard|awk '{print $2}'`

`export GIN_MODE=release`

if [ "$_pid" == "" ]
then
    bin=`dirname $0`
    root=`(cd $bin/../ && pwd)`
    cd $root
    mkdir -p logs
    `govendor build`
    nohup $root/gin-msgboard -conf $conf > $log 2>&1 &
    sleep 3
    _pid=`ps -ef | grep -v grep|grep gin-msgboard|awk '{print $2}'`
    if [ "$_pid" != "" ]
    then
        echo "Started, pid: $_pid"
         sleep 3
    else
        echo "Failed to start"
        cat $log
    fi
else
    echo "Error: already running with pid: $_pid"
    exit 1
fi
