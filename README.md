# FaaS Examples

A collection of serverless functions for FaaS experimentation.

## Available Functions

| Function                                   | Description                                                                     | Runtime        |
| ------------------------------------------ | ------------------------------------------------------------------------------- | -------------- |
| [Hello World](./hello-world)               | A simple function that exposes an HTTP endpoint and returns a greeting message. | Python         |
| [Image Resizer](./image-resizer)           | A function that resizes an image to a specified width and height.               | Python         |
| [Matrix Multiplier](./matrix-multiplier)   | A function that multiplies two matrices.                                        | Python (NumPy) |
| [Vector Magnitude](./vector-magnitude)     | A function that calculates the magnitude of a vector.                           | Python (NumPy) |
| [Weather Forecast](./weather-forecast)     | A function that retrieves the weather forecast for a given location.            | Go             |
| [Palindrome Checker](./palindrome-checker) | A function that checks if a given string is a palindrome.                       | Go             |
| [Word Counter](./word-counter)             | A function that counts words in a provided text.                                | Go             |
| [UUID Generator](./uuid-generator)         | A function that generates and returns a UUID.                                   | Go             |
| [Fibonacci](./fibonacci)                   | A function that generates the Fibonacci sequence up to a specified number.      | Node.js        |
| [Spring Boot](./spring-boot)               | A simple Spring Boot application that exposes an HTTP endpoint.                 | Java           |

## Pre-built Docker Images

Pre-built Docker images are available on Docker Hub:

- [devandrejesus/hello-world-python](https://hub.docker.com/repository/docker/devandrejesus/hello-world-python)
- [devandrejesus/image-resizer-python](https://hub.docker.com/repository/docker/devandrejesus/image-resizer-python)
- [devandrejesus/matrix-multiplier-python](https://hub.docker.com/repository/docker/devandrejesus/matrix-multiplier-python)
- [devandrejesus/vector-magnitude-python](https://hub.docker.com/repository/docker/devandrejesus/vector-magnitude-python)
- [devandrejesus/weather-forecast-go](https://hub.docker.com/repository/docker/devandrejesus/weather-forecast-go)
- [devandrejesus/palindrome-checker-go](https://hub.docker.com/repository/docker/devandrejesus/palindrome-checker-go)
- [devandrejesus/word-counter-go](https://hub.docker.com/repository/docker/devandrejesus/word-counter-go)
- [devandrejesus/uuid-generator-go](https://hub.docker.com/repository/docker/devandrejesus/uuid-generator-go)
- [devandrejesus/fibonacci-node](https://hub.docker.com/repository/docker/devandrejesus/fibonacci-node)
- [devandrejesus/spring-boot-java](https://hub.docker.com/repository/docker/devandrejesus/spring-boot-java)

Since my goal is to test the [stargz snapshotter](https://github.com/containerd/stargz-snapshotter), I built these images using the estargz compression format. I used the following command to build them:

```bash
docker buildx build --tag devandrejesus/uuid-generator-go:esgz -f .\Dockerfile -o type=registry,oci-mediatypes=true,compression=estargz,force-compression=true .
```

But they can be built without the estargz compression format.
