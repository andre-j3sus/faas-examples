# Fibonacci

A function that generates the Fibonacci sequence up to a specified number.

## Build

```bash
docker build -t faas-fibonacci .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name faas-fibonacci faas-fibonacci
```

## Example Call

```bash
# Call the function with custom limit
curl http://localhost:8080/?limit=15

# Expected response for limit=15:
# {"status":"success","count":15,"sequence":[0,1,1,2,3,5,8,13,21,34,55,89,144,233,377]}
```
