# Palindrome Checker

A function that checks if a given string is a palindrome.

## Build

```bash
docker build -t faas-palindrome-checker .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name faas-palindrome-checker faas-palindrome-checker
```

## Example Call

```bash
# Check if a string is a palindrome
curl -X POST \
  -H "Content-Type: application/json" \
  -d '{
    "text": "A man, a plan, a canal: Panama"
  }' \
  http://localhost:8080/check

# Expected response:
# {
#   "text": "A man, a plan, a canal: Panama",
#   "isPalindrome": true,
#   "explanation": "The text reads the same backward as forward (ignoring case, spaces, and non-alphanumeric characters)."
# }
```
