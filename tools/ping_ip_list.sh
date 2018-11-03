#!/bin/bash
# This script will ping a list of IP addresses and wait for a response
# If they didn't reply in one second, the IP is considered dead.

seconds=1
filename="ip_uniques.txt"

echo "You can't CTRL^C this script once it's started."
read -p $'Press enter to continue\n'

while read p; do
	ping $p -c 1 -w $seconds  >/dev/null && echo "$p"
done <$filename
