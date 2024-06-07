# FilterID IP Address API

This application is a simple API that returns the requester's IP address in JSON, plain text, or XML format.

## Table of Contents
- [Building the App](#building-the-app)
- [Using the Makefile](#using-the-makefile)
- [Using Docker Compose](#using-docker-compose)
- [Security Considerations](#security-considerations)
- [IP Address Retrieval](#ip-address-retrieval)
- [Building Tailwind CSS](#building-tailwind-css)

## Building the App

To build the application, you need to have Go installed. Make sure you have Go tools set up properly. You can download Go from [golang.org](https://golang.org/dl/).

Once Go is installed, navigate to the project directory and run:

```bash
go build
```

This command will compile the application and create an executable file in the project directory.

## Using the Makefile

For convenience, you can also use the provided Makefile to build the application. Simply run:

```bash
make
```

This will execute the build commands defined in the Makefile.

## Using Docker Compose

To build and run the application using Docker Compose, you can use the following command:

```bash
docker compose build
```

This will build the Docker images defined in `compose.yml`. You can start the application with:

```bash
docker compose up
```

### Docker Compose File Explanation

The `compose.yml` file defines the services, networks, and volumes for your application. It specifies how to build the Docker images and any configurations required for the containers. Make sure to review the file for any service-specific settings.

## Security Considerations

For security purposes, it's recommended to add a reverse proxy in front of your API if necessary. A reverse proxy can help log requests and provide an additional layer of security. Also, consider implementing rate limiting to prevent abuse of the API.

## IP Address Retrieval

In this application, we determine the requester's IP address by checking various headers. We prioritize the following in order:

1. **Cloudflare Header**: If you're using Cloudflare, we look for the `CF-Connecting-IP` header.
2. **Proxy Headers**: We check for `X-Forwarded-For` and similar headers to get the original IP if the request is coming through a proxy.
3. **Direct IP**: If no headers are found, we fall back to using the direct IP address from the request.

This method ensures that we get the most accurate IP address possible.

## Building Tailwind CSS

To build the Tailwind CSS for the application, you have several options:

1. **Using the Makefile**:
   To build Tailwind CSS, you can use the Make command:
   ```bash
   make build-tailwind
   ```
   This utilizes Docker Compose to build the CSS files in the Tailwind builder service.

2. **Using the `build-tailwind.sh` script**:
   You can also run the provided shell script to build Tailwind CSS:
   ```bash
   ./build-tailwind.sh
   ```
   This script uses the Docker Node container to execute the necessary commands for compiling Tailwind CSS, providing a convenient alternative.
   
3. **Using npm**:
   You can build it directly with npm/node:
   ```bash
   npm install
   npm run build
   ```
   Make sure you have Node.js and npm installed. The Tailwind CSS is configured to style the application, and the build process generates the necessary CSS files.