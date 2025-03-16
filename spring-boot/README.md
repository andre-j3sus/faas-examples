# Spring Boot Application

A containerized Spring Boot application that provides RESTful API endpoints.

This example is from [spring-guides/gs-spring-boot](https://github.com/spring-guides/gs-spring-boot).

## Build

```bash
docker build -t spring-boot-app .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name spring-boot-app spring-boot-app
```

## Example Call

```bash
curl http://localhost:8080/

# Expected response:
# {"message":"Hello, World!","timestamp":"2025-03-16T12:34:56.789Z"}
```
