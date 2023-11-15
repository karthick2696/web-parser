# Web Page Fetcher in Go

This is a simple command-line program implemented in Go that fetches web pages and saves them to disk for later retrieval and browsing.

## Getting Started

### Prerequisites

Make sure you have docker installed on your machine.

### Cloning the Repository

Clone the repository:

```bash
git clone https://github.com/karthick2696/web-parser
```

Changing Directory (cd) to the Repository:

```bash
cd web-parser
```

### Building the Program

Build the program

```bash
docker build -t web-parser .
```

### Running the Program

To fetch a web page, run the following command

#### Run program without metadata

```bash
docker run web-parser https://www.google.com https://autify.com
```

#### Run program with metadata:

```bash
docker run web-parser --metadata https://www.google.com https://autify.com
```
 Output :
 
 <img width="920" alt="image" src="https://github.com/karthick2696/web-parser/assets/30376444/55d8dafc-3b4f-49be-babe-ca2e79f32f55">

 ### To View HTML file 

 We are running program using docker so html file of website will be stored in docker volumes. Due to that we need to move html files from docker to local  

 #### Step 1: 
 
 Get docker container id of web-parser using below commend 

```bash
docker ps -a
```
Output :

Copy container id 

<img width="1024" alt="image" src="https://github.com/karthick2696/web-parser/assets/30376444/7733da6b-91bc-4e87-ad12-b826a200b811">

#### Step 2:

Copy html files from docker container based on container-id using below commend 

```bash
docker cp container-id:/app/. ./output
```
Replace container-id by value we copied at first step 

Output :

Files will be moved to output folder of your local code repo and just nevigate to output folder to view html files like below in image 

<img width="886" alt="image" src="https://github.com/karthick2696/web-parser/assets/30376444/cb471b42-9294-49c7-b571-f578c61f195b">

### Important Condition :- 

When ever we run program using docker new container-id will be generated each time. Due to that we need take latest container id to copy files from docker to local machine. otherwise it will copy old files only from old container 
