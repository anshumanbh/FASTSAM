FROM ubuntu:14.04
MAINTAINER Anshuman Bhartiya

RUN apt-get update
RUN apt-get install -y git build-essential curl wget libpcap-dev libdigest-hmac-perl libdigest-md5-file-perl libfindbin-libs-perl libmime-base64-urlsafe-perl libgetopt-long-descriptive-perl && apt-get clean 

RUN git clone https://github.com/robertdavidgraham/masscan /opt/masscan
WORKDIR /opt/masscan

RUN make -j

RUN cp /opt/masscan/bin/masscan /usr/local/bin

RUN mkdir /opt/secdevops
COPY masscan/scripts/* /opt/secdevops/
RUN chmod +x /opt/secdevops/*

# Enable these once you have s3curl configured to upload the scan results to your S3 object store.
# You would basically need 2 files - .s3curl (creds) and s3curl.pl (script to upload)

# RUN mkdir /opt/s3curl
# COPY s3curl/* /opt/s3curl/
# RUN chmod 0400 /opt/s3curl/.s3curl

ENV PROJECT 123456

CMD ["/opt/secdevops/masscan.sh"]
