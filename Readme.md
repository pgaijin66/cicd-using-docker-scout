# cicd-using-docker-scout

This project demonstrates Continuous Integration and Continuous Deployment (CI/CD) using Docker and Scout.

## Build

To build the Docker image for your application, use the following command:

```bash
docker build -t my-app:buggy .
```

## Running application

To build the Docker image for your application, use the following command:

```bash
docker run -p 9090:9090 my-app:buggy
```

## Scanning using Docker Scout

To build the Docker image for your application, use the following command:

```bash
docker scout cves my-app:buggy
```



