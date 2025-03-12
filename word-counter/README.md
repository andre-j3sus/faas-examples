# Word Counter

A function that counts words in a provided text.

## Build

```bash
docker build -t faas-word-counter .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name faas-word-counter faas-word-counter
```

## Example Call

```bash
# Count words in a text
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "text": "The quick brown fox jumps over the lazy dog. The dog was too lazy to react."
  }' \
  http://localhost:8080/count

# Expected response:
# {
#   "text": "The quick brown fox jumps over the lazy dog. The dog was too lazy to react.",
#   "wordCount": 16,
#   "charCount": 74,
#   "frequentWords": {
#     "the": 3,
#     "lazy": 2,
#     "dog": 2,
#     "quick": 1,
#     "brown": 1
#   }
# }
```
