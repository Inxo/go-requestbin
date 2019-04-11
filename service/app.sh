#!/bin/sh
### BEGIN INIT INFO
# Provides:          GoRequestBin
# Required-Start:
# Required-Stop:
# Default-Start:     2 3 4 5
# Default-Stop:      1 0 6
# Short-Description: ...
# Description: ...
### END INIT INFO
case $1 in
    start)
        echo "Starting requestbin web app."
	cd /home/inxo/gorequestbin/
        sudo -u inxo /home/inxo/gorequestbin/./app >> /home/inxo/gorequestbin/log/requestbin.log &
        ;;
    stop)
        echo "Stopping requestbin web app."
        sudo kill $(sudo lsof -t -i:8004)
        ;;
    *)
        echo "Requestbin web app service."
        echo $"Usage $0 {start|stop}"
        exit 1
esac
exit 0