echo "You can't CTRL^C this script once it's started."
read -p $'Press enter to continue\n'

while read p; do
	ping $p -c 1 -w 1  >/dev/null && echo "$p"
done <ip_uniques.txt
