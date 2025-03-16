# File Manager

A function that exposes an API for managing files in a specified directory.

## Build

```bash
docker build -t faas-file-manager .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name faas-file-manager faas-file-manager
```

## Example Call

```bash
# Call the function (assumes it's running on port 8080)
curl http://localhost:8080/read?path=/tmp/test.txt

# Call the function with a POST request
curl -X POST -d '{"path": "/tmp/test.txt", "content": "Hello, World!"}' http://localhost:8080/write

# Call the function to list files in a directory
curl http://localhost:8080/list?path=/tmp
```
