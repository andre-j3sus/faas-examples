# Image Resizer

A function that takes an uploaded image and resizes it to the dimensions specified in the request parameters. It returns the resized image as a downloadable file.

## Build

```bash
docker build -t faas-image-resizer .
```

## Run

```bash
# Run the container and expose it on port 8080
docker run -d -p 8080:8080 --name faas-image-resizer faas-image-resizer
```

## Example Call

```bash
# Resize an image to 300x200 pixels
curl -X POST \
  -F "image=@path/to/your/image.jpg" \
  -F "width=300" \
  -F "height=200" \
  http://localhost:8080/resize \
  --output resized_image.png
```
