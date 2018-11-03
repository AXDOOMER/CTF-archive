#!/bin/bash
# Removes duplicates entries and sorts them
# Great for a list of IP addresses

echo "Paste a bunch of text in here to remove identical lines and sort them."
echo "DO NOT paste anything that contains an empty line, "
echo "  or they will EXECUTE in your terminal!"

ARRAY=()

while read -r el; do
	ARRAY+=("$el")

	if [ -z "$el" ]; then
		break
	fi
done

IFS=$'\n' ARRAY=($(sort <<<"${ARRAY[*]}" | uniq))

for el in "${ARRAY[@]}"; do
	echo $el
done
