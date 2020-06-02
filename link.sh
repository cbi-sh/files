##!/bin/bash
#
#host="cbi-test1"
#user="dev"
#
#source="."
#destination="/store/app/src/cbi"
#echo "link to host ${host} established"
#
#fswatch -9 -e ".*" -i "\.go$" . | \
#xargs   -9 -n 9 -I {} #\
##rsync   -aruz --include="*.go" --exclude="*" --delete --progress ${source} ${user}@${host}:${destination}
