# UUID Generator

A function that generates and returns a UUID.

## Build

```bash
docker build -t faas-uuid-generator .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name faas-uuid-generator faas-uuid-generator
```

## Example Calls

### Generate a single UUID

```bash
# Generate a single UUID
curl http://localhost:8080/uuid

# Expected response:
# {
#   "uuid": "f47ac10b-58cc-4372-a567-0e02b2c3d479",
#   "version": 4,
#   "timestamp": "2025-03-12T15:04:05Z"
# }
```

### Generate a batch of UUIDs

```bash
# Generate multiple UUIDs
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "count": 3
  }' \
  http://localhost:8080/batch

# Expected response:
# {
#   "uuids": [
#     "f47ac10b-58cc-4372-a567-0e02b2c3d479",
#     "5b6c8d3f-21a4-44e7-9cb5-f8d6e7c5b4a3",
#     "9e8d7c6b-5a4f-43e2-b1c0-d9e8f7c6b5a4"
#   ],
#   "count": 3,
#   "timestamp": "2025-03-12T15:04:05Z"
# }
```
