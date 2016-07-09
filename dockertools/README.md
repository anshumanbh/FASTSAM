dockertools
============

Security Tools that can be started as individual docker containers 

* Make sure you have the .s3curl file with creds in it inside the s3curl folder and the s3curl.pl script to be able to upload the scan/test results to the S3 bucket. Since I am actively developing this at my work place, I can't unfortunately paste our scripts and creds here. I will, however, come back to this once I have the major parts of this framework developed.

* You have a choice of either building all tools via the `docker-compose.yml` file by simply issuing `docker-compose build` from the `dockertools` directory OR building individual tools as you wish.

* If you don't want to build all the tools in here and if you only want to build specific tools for your environment, you would navigate to the `dockertools` directory and issue the following commands respectively:
	
	* `docker build -t dockertools_nmap -f nmap/Dockerfile .` for building the nmap tool
	* `docker build -t dockertools_masscan -f masscan/Dockerfile .` for building the masscan tool

* The end result is the same that there were will be 2 images built - `dockertools_nmap` and `dockertools_masscan` in your docker environment.

NOTE: Remember that the names should be exactly like the ones mentioned above because those are the names used by the API.

* For my environment and the architecture described, I use `docker-machine` on my local workstation and I have configured it to talk to a remote docker host. So, whatever I do in terms of docker on my local workstation is actually being performed on my remote docker host. If you need help with this, please message me. It is very easy and I can help you get started here as well. As always, I will come back to this once I am done with the higher priority tasks. You can very well treat your local workstation as the docker host as well if you don't want to configure a remote docker host. 





