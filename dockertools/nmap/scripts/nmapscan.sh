#!/usr/bin/env bash
nmap --script-updatedb

# Running TCP Scan
nmap -A -sS -sC -sV -vv -Pn -p- -T 4 -oX /tmp/tcpscan.xml -oN /tmp/tcpscan.nmap $TARGETS

# Running UDP Scan. This is taking way too long. Won't implement UDP scan just yet
# nmap -A -sU -sC -sV -vv -Pn --top-ports 200 -T 4 -oX /tmp/udpscan.xml -oN /tmp/udpscan.nmap $TARGETS

# Sending both tcpscan and udpscan nmap outputs to S3. 
# Enable this once you have s3curl configured.
# perl -w /opt/s3curl/s3curl.pl --put=/tmp/tcpscan.nmap -- https://<domain>/$PROJECT/nmap/tcpscan.nmap
# perl -w /opt/s3curl/s3curl.pl --put=/tmp/udpscan.nmap -- https://<domain>/$PROJECT/nmap/udpscan.nmap
