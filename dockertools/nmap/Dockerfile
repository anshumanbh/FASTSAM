FROM kalilinux/kali-linux-docker
MAINTAINER Anshuman Bhartiya

RUN echo "deb http://http.kali.org/kali kali-rolling main contrib non-free" > /etc/apt/sources.list && \
echo "deb-src http://http.kali.org/kali kali-rolling main contrib non-free" >> /etc/apt/sources.list
ENV DEBIAN_FRONTEND noninteractive
RUN apt-get -y update && apt-get -y dist-upgrade && apt-get clean

RUN apt-get -y install nmap curl libdigest-hmac-perl libdigest-md5-file-perl libfindbin-libs-perl libmime-base64-urlsafe-perl libgetopt-long-descriptive-perl && apt-get clean

RUN mkdir /opt/secdevops
COPY nmap/scripts/* /opt/secdevops/
RUN chmod +x /opt/secdevops/*

# Enable these once you have s3curl configured to upload the scan results to your S3 object store.
# You would basically need 2 files - .s3curl (creds) and s3curl.pl (script to upload)

# RUN mkdir /opt/s3curl
# COPY s3curl/* /opt/s3curl/
# RUN chmod 0400 /opt/s3curl/.s3curl

ENV PROJECT 123456

CMD ["/opt/secdevops/nmapscan.sh"]
