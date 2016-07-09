#!/usr/bin/env bash

IFS=' ' read -r -a array <<< "$TARGETS"

for ip in "${array[@]}";
do
		echo $ip >> /tmp/iplist.txt;
done

# Running a Scan
masscan --ports 0-65535 --banners -oX /tmp/masscan.xml --rate 1000000 --includefile /tmp/iplist.txt

# Sending the scan output to S3
# Enable this once you have s3curl configured.
# perl -w /opt/s3curl/s3curl.pl --put=/tmp/masscan.xml -- https://<domain>/$PROJECT/masscan/masscan.xml
