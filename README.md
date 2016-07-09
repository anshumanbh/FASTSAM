# FASTSAM
Framework for Automated Security Testing that is Scaleable and Asynchronous built on Microservices 

![GitHub Logo](/images/Architecture.png)

## DISCLAIMER
This is just an idea of how I envision this framework to be built from the ground up. It is still very much a work in progress. I have already started building this with a few key elements up and running. To start with, I am focussing on just automating some basic scanning using nmap and masscan. 

## OVERVIEW
Based on the above architecture diagram, the framework would do the following:

1. Take input from an end user from a UI. This input will be IPs, IP ranges, domains, etc. 

2. Submit a request to an API server with the above input parameters to specific API endpoints. Some examples are:

	. POST /api/v1/nmapscan
	. POST /api/v1/masscan

3. This API server would then submit a task in an asynchronous message processing queue. I have used [Machinery](https://github.com/RichardKnop/machinery) for this that uses `Rabbitmq` as the broker/backend and a worker to pick up tasks from the queue and process them asynchronously. 

4. Whenever a task is submitted, a task ID is returned which can then be used by the UI to show the status and progress of that particular task, if needed.

5. The API server is also connected to a remote Docker Host (via docker-machine) that has a docker engine running. This docker host basically has all the microservies running on it as Docker containers. 

6. The workers pick up a task from the queue and start docker containers on this remote docker host to process them. So, for example, if a nmap task is submitted, the worker would start a nmap docker container in the docker host that would run nmap against the target IPs/domains.

7. Once the scan finishes, the results are uploaded to an object store like S3 and the Docker containers are shut down and removed not leaving any mess behind and with artifacts properly stored.

8. The API server then connects to the S3 object store and uploads the scan artifacts to be visualized. I am using LAIR (and its drones) for this since I feel its a nice UI for triaging. But, other UI options can also be explored such as Dradis, ElasticSearch, EyeWitness?, etc.    

That's it for now! 

## Key features of this framework
* It uses security tools as Docker containers. You can build as many tools as you want to and use them as per your requirements.
* It uses asynchronous processing of tasks that gives more flexibility of running multiple scans by multiple tools on multiple targets.
* It separates the tooling logic from the main API logic so tools can be built separately and the API can be as light as it can be.
* It uses Docker which is awesome when you want to run quick tools and destroy the environment when the job is done.