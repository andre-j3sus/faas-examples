# Matrix Multiplier

A function that takes two matrices as input and returns their product. It uses NumPy for efficient matrix operations and validates that the matrices can be multiplied together.

## Build

```bash
docker build -t faas-matrix-multiplier .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name faas-matrix-multiplier faas-matrix-multiplier
```

## Example Call

```bash
# Multiply two 2x2 matrices
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "matrix_a": [[1, 2], [3, 4]],
    "matrix_b": [[5, 6], [7, 8]]
  }' \
  http://localhost:8080/multiply

# Expected response:
# {
#   "result": [[19.0, 22.0], [43.0, 50.0]],
#   "shape": {"rows": 2, "columns": 2}
# }
```
