# Yahtzee
## Prompt
PLEASE READ THIS ENTIRE README AND THE ENTIRE OPENAPI DOCUMENT BEFORE YOU BEGIN WRITING TEST CASES. Don't forget to read the description at the top of the OpenAPI document - it contains information that will be very helpful for your task.

This repository contains the source code for a new service we are shipping called "Yahtzee". You are tasked with writing test cases for the Yahtzee API to ensure it behaves as documented in the OpenAPI document. Over the course of the next ~3 hours, in the language and testing framework of your choice, please write the necessary test cases to validate the Yahtzee service. If you run out of time before you can implement all of your test cases, please add comments and/or pseudocode for test cases you planned to write but did not have time to complete.

IMPORTANT: You are not required to run the Yahtzee service in order to complete this assignment. If you are unfamiliar with both Docker and Go or you do not have an accessible environment where these tools are readily available, we recommend you skip trying to run the Yahtzee service and focus on your test plan and test cases. Your test plan and the test cases you write for this exercise are the most important part. You should write your test cases based on the contents of the OpenAPI document. When you meet with the interview panel, please inform them whether or not you were able to execute your test cases against a running instance of the Yahtzee service. You will not be penalized if you have not executed your test cases against a live version of the Yahtzee service.
## Repository contents
### OpenAPI documents
The OpenAPI specification is provided in YAML format. Feel free to import openapi.yaml into a Swagger/OpenAPI editor to view the rendered version. The description at the top of the document contains helpful information such as the expectations of the Yahtzee service and the username and password used to authenticate with the Yahtzee service.
### Dockerfile
A Dockerfile which can be used to build the Yahtzee service into a container is provided. Details on how to do this are provided below.
### Source code
The source code (the *.go files) for the Yahtzee service are provided. You are welcome to browse the source code, but this should not be required to complete this task. You are also welcome to build and run the source code.
## Running the service
By default, the Yahtzee service listens on port 8080. This can be modified if 8080 is not available on your system.
### Docker (preferred method)
You must have an installation of Docker available to build and run the service via Docker.
#### Build the container

    <clone the repository>
    cd /path/to/yahtzee
    docker build -t yahtzee:latest .
    
#### Run the container

    docker run -it -p 8080:8080 yahtzee:latest
#### Run the container on a different port
To bind the Yahtzee service to a different port, use the following command:

    docker run -it -p <new-port>:8080 yahtzee:latest
    docker run -it -p 8081:8080 yahtzee:latest

### Go
If you have a Go installation available, you can compile and run this service directly. This service was built using Go 1.18.4. Other versions of Go may work, but they have not been tested.
#### Run the service

    <clone the repository>
    cd /path/to/yahtzee
    go run .
 
#### Run the service on a different port
The Yahtzee service accepts a parameter "-port" which may used to change the port number it listens on:

    go run . -port <new-port>
    go run . -port 8081
