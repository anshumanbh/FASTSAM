# FASTSAM
### Framework for Automated Security Testing that is Scaleable and Asynchronous built on Microservices 

![Architecture Logo](/images/Architecture.png)

## DISCLAIMER
This is just an idea of how I envision this framework to be built from the ground up. It is still very much a work in progress. I have already started building this with a few key elements up and running. To start with, I am focussing on just automating some basic scanning using nmap and masscan. 

## OVERVIEW
Based on the above architecture diagram, the framework would do the following:

1. Take input from an end user from a UI. This input will be IPs, IP ranges, domains, etc. 

2. Submit a request to an API server with the above input parameters to specific API endpoints. Some examples are:

		* POST /api/v1/nmapscan

		* POST /api/v1/masscan

3. This API server would then submit a task in an asynchronous message processing queue. I have used [Machinery](https://github.com/RichardKnop/machinery) for this that uses `Rabbitmq` as the broker/backend and a worker to pick up tasks from the queue and process them asynchronously. 

4. Whenever a task is submitted, a task ID is returned which can then be used by the UI to show the status and progress of that particular task, if needed.

5. The API server is also connected to a remote Docker Host (via docker-machine) that has a docker engine running. This docker host basically has all the microservies running on it as Docker containers. 

6. The workers pick up a task from the queue and start docker containers on this remote docker host to process them. So, for example, if a nmap task is submitted, the worker would start a nmap docker container in the docker host that would run nmap against the target IPs/domains.

7. Once the scan finishes, the results are uploaded to an object store like S3 and the Docker containers are shut down and removed not leaving any mess behind and with artifacts properly stored.

8. The API server then connects to the S3 object store and uploads the scan artifacts to be visualized. I am using LAIR (and its drones) for this since I feel its a nice UI for triaging. But, other UI options can also be explored such as Dradis, ElasticSearch, EyeWitness?, etc.    

That's it for now! 

## Key features of this framework
* API-like framework to perform security testing/scanning activities.
* Asynchronous processing of tasks that gives more flexibility of running multiple scans by multiple tools on multiple targets.
* Ability to query the task status using the task ID of a job.
* Remote Dockerized environment for microservices that can be anything from security tools to visualization tools to triaging tools. You can build as many tools as you want to and use them as per your requirements.
* Separate tooling logic from the main API logic so tools can be built separately and the API can be as light as it can be.
* Flexibility of Docker to quickly spawn up and destroy environments.
* Uploading artifacts to an object store.
* Using existing open source tools and integrating them together without reinventing the wheel.

## Let's get started
* To begin with, git clone this repo.

* Navigate to the `dockertools` directory. This directory will have all the tools that we would need to perform your security related tasks. There is also a `README` in this directory that describes what you would need to do in order to build all tools together or build specific tools only.

*  So, go build those tools first. If you want to try out both nmap and masscan, first ensure you have docker and docker-compose setup, then simply run `docker-compose build` after navigating to the `dockertools` directory. You should see `dockertools_nmap` and `dockertools_masscan` images built in your environment when you type `docker images`.

* If you just build the tools as is, it *should* build just fine. However, the scan reports won't be sent to S3 because obviously you would need to configure that according to your environment. Right now, I have commented out all the parts in the `nmap/Dockerfile`, `masscan/Dockerfile`,  `nmap/scripts/nmapscan.sh` and `masscan/scripts/masscan.sh` files that need to be uncommented if the S3 integration has to work. But, before you uncomment those parts, make sure you update the `s3curl/s3curl.pl` script to add the logic to upload the results to S3 and also provide the S3 creds in the `s3curl/.s3curl` file.

* Once you have the above docker images ready for the nmap and masscan tools, we can start the API server. Remember, the flow that we want to achieve is: 

```API request from a UI or CURL request -> API Server -> Machinery (Rabbitmq -> Worker) -> Use the above Docker images to start containers on the Docker host -> Containers spawn up, scan, upload the scan result to S3 (won't work out of the box unless you configure your S3 environment) -> Containers get destroyed -> scan results from S3 uploaded to a visualization tool (work in progress)```

* You can start the API server remotely but since this is still very much a work in progress, it makes more sense to start the API server locally and test it without having to do it remotely. The first thing you need here is `rabbitmq-server`installed and running on port `5672` locally acting both as broker and result backend. You should be able to connect to it via `amqp://guest:guest@localhost:5672/`. I used `homebrew` on my mac to install this. To start the rabbitmq server, just type `rabbitmq-server`. You should see something like below:

![Rabbitmq Logo](/images/rabbitmq.png)

* In order to proceed further, you also need `GO` installed and configured on your workstation. After installing `rabbitmq`, run the `runworker.go` file (in a separate terminal) by typing `go run runworker.go` from the runworker directory in a terminal. This will start a worker connected to the Rabbitmq server as the broker and backend. In this runworker.go file, please take care of the `tasks` import statement depending upon how your GO environment has been configured. I have left a comment there to explain it better. See pic below:

![Runworker Logo](/images/runworker.png)

* Now, run all the remaining *.go files from the `api` directory by navigating to it and typing `go run *.go` (in a separate terminal). This will start our `main.go` file with the router and the API running on port `8080` locally. 

* At this point, you are ready to start firing the API requests to the API server. So, in yet another terminal, send a CURL request to our GO API by typing the following CURL commands:

	* `curl -H "Content-Type: application/json" -d '["127.0.0.1","scanme.nmap.org"]' http://localhost:8080/api/v1/nmapscan`.
	* `curl -H "Content-Type: application/json" -d '["127.0.0.1","10.0.0.0/8"]' http://localhost:8080/api/v1/masscan`.

Remember nmap takes IPs, IP ranges, domains, etc. separated by a whitespace whereas Masscan only takes IPs and IP ranges, that too in a file.

* These requests will immediately submit tasks to the worker (of running nmap and masscan against all these IPs/domains) and return back the task ID which can be used to later query the status of the task. Its a non-blocking asynchronous event. It is important to note here that the endpoint /api/v1/nmapscan and /api/v1/masscan submits these tasks. These tasks (when run) will take the environment of the docker-machine of the terminal where these tasks were called from. So, if you don't have the `dockertools_nmap` and `dockertools_masscan` images running in that docker environment, you might have to build it first. Refer to the dockertools README in order to get more details about building these images. These tasks will be run in the background after getting picked by the worker. Once the tasks finishes, the results will be returned to the worker.

## TODO

- [ ] Write more scanners/tools Dockerfiles and push it to the docker lab so that we can just call those scanners and get our results.

- [ ] Write scanners for websites, subdomain enumerators, etc.

- [ ] The LAIR framework, their drones and even their parsers are written in GO. So, figure out if we can leverage any of that code or use their API to depict all the info in a UI like LAIR does after running our scanners/tools. Also, worth exploring if any of the scan stuff can be depicted in Dradis or Eyewitness or rawr

- [ ] Automate portions after the initial scanning. This is where the main benefit is i.e. automate tasks for pentesting. Explore Gauntlt for this phase for automation test cases.

- [ ] Make modifications for reading IP ranges, XML, text file uploads of IPs and domains to be scanned.

- [ ] Right now, the scans are uploaded in the S3 bucket as a common name. Figure out a way to append something to the filename to be able to distinguish what hosts were scanned when. 

- [ ] Implement authentication and integrate this with a UI frontend. Also, process background AJAX calls to check the status of the tasks submitted to machinery.

- [ ] Change the logic of docker remote API where instead of taking the config from the docker-machine environment of a terminal, it can be provided by a file so that we don't have to rely on the docker-machine env.

- [ ] Move the amqp settings from the code to a .yml file.

- [ ] Figure out a way to keep all the tools/repos uptodate via some docker update mechanism or a cron job of daily pulling fresh images.

- [ ] Writes test cases.

- [ ] Make this work against a general S3 environment.

- [ ] Add instructions on starting a remote docker host from your local workstation using docker-machine.

