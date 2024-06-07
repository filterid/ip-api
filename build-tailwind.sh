#!/bin/bash

# Run the tailwind build command inside the container
docker run --rm \
    -v $(pwd):/app \
    -w /app node:18-alpine \
    sh -c "npm install && npx tailwindcss -i ./app.css -o ./static/css/tailwind.css --minify"
