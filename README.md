# Go REST Authentication Base (go-rest-auth-base)

Template for creating REST APIs in Go with built-in authentication middleware and logging capabilities.
This project is a starting point for those who want to explore these concepts in Go. Remember, it's a template - use it to learn, but consider more robust solutions for production applications.

## Features

- API Key Authentication: Incoming requests are authenticated by validating the 'X-Api-Key' present in the request headers.
- Rate Limiting: Option to limit the number of requests per API key for a given duration.
- Thread-Safe API Key Mapping: Thread-safe map to store and manage API key objects.
- Request Logging: Log details about each incoming request.

## Usage

```bash
$ git clone https://github.com/fxnatic/go-rest-auth-base.git
$ cd go-rest-auth-base
$ go run main.go
```

## License

This project is licensed under the [MIT License](/LICENSE).