# Vector Magnitude

A function that calculates the Euclidean norm (magnitude) of a vector. It accepts a vector as a JSON array and returns its magnitude along with the original vector and its dimensions.

## Build

```bash
docker build -t faas-vector-magnitude .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name faas-vector-magnitude faas-vector-magnitude
```

## Example Call

```bash
# Calculate the magnitude of a 3D vector [3, 4, 5]
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "vector": [3, 4, 5]
  }' \
  http://localhost:8080/magnitude

# Expected response:
# {
#   "vector": [3.0, 4.0, 5.0],
#   "magnitude": 7.0710678118654755,
#   "dimensions": 3
# }
```
