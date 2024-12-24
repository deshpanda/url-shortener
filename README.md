# url-shortener

This is a simple URL shortener tool that allows you to shorten long URLs and redirect to the original URLs using the generated short code.

## Features:
- Generate short URLs.
- Redirect to the original URL using the shortened URL.

## Steps to Use

### 1. Clone the Repository
First, clone the repository to your local machine if you haven't done so already.

```bash
git clone <repository-url>
cd url-shortener
```

### 2. Run the application
Run the Go server to start the URL shortener service.

```bash
go run main.go
```
This will start an HTTP server on `http://localhost:8080`.

### 3. Shorten a URL
To shorten a URL, send a `POST` request to the `/shorten` endpoint with the long URL as the body parameter.

Use curl to send the request:

```bash
curl -X POST -d "url=https://x.com" http://localhost:8080/shorten
```

### 4. View the Shortened URL
The server will return a short URL, which you can access to be redirected to the original URL. For example:

```bash
Short URL: http://localhost:8080/sRDJAH
```
You can now visit `http://localhost:8080/sRDJAH`, and it will redirect to `https://x.com`.

### 5. Access the Shortened URL
You can also use `curl` to access the shortened URL:

```bash
curl http://localhost:8080/sRDJAH
```
This will redirect to the original URL (`https://x.com`).

### Sample Response
After sending the `POST` request to shorten the URL, the output will look like this:

```bash
Short URL: http://localhost:8080/sRDJAH
```
If you visit the shortened URL (`http://localhost:8080/sRDJAH`), it will automatically redirect you to the original URL.

### Error Handling
If the provided URL is not valid, or the request method is incorrect, you will receive an appropriate error message.