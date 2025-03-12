# Hello World

A function that exposes a basic HTTP endpoint that returns a JSON response with a greeting message.

## Build

```bash
docker build -t faas-hello-world .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name faas-hello-world faas-hello-world
```

## Example Call

```bash
# Call the function (assumes it's running on port 8080)
curl http://localhost:8080/

# Expected response:
# {"message":"Hello, World!","status":"success"}
```
