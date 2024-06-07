# Build Go application
docker-build:
	docker-compose build

# Start the Docker container
docker-up:
	docker-compose up -d

# Stop the Docker container
docker-down:
	docker-compose down

# Build Tailwind CSS using Docker
build-tailwind:
	docker-compose run --rm tailwind-builder npx tailwindcss -i ./static/css/input.css -o ./static/css/styles.css --minify

# Run the Go app without Docker (for local testing)
run:
	go run main.go
