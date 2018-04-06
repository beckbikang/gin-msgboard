#!/bin/bash
_pid=`ps -ef | grep -v grep|grep gin-msgboard|awk '{print $2}'`

if [ "$_pid" == "" ]
then
    echo "Error: not running"
    exit 1
else
    kill $_pid
    if [ $? -ne 0 ]
    then
        echo "Error: failed to kill $_pid"
        exit 2
    else
        echo "Stopped"
    fi
fi