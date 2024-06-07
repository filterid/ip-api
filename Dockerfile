# Step 1: Build Tailwind CSS

FROM node:18-alpine AS tailwind-builder
WORKDIR /app
COPY ./tailwind.config.js ./package.json ./postcss.config.js ./
COPY ./static/css/input.css ./static/css/
RUN npm install
RUN npx tailwindcss -i ./static/css/input.css -o ./static/css/styles.css --minify

# Step 2: Build Go app

FROM golang:1.22.4-alpine AS go-builder
WORKDIR /app
COPY go.mod go.sum ./
RUN go mod download
COPY . .
RUN go build -o /filterid-ip-api .

# Step 3: Final stage with built Go app and Tailwind CSS
FROM alpine:latest
WORKDIR /app
COPY --from=go-builder /filterid-ip-api /app/filterid-ip-api
COPY --from=tailwind-builder /app/static/css/styles.css /app/static/css/styles.css
COPY ./static /app/static
COPY ./templates /app/templates
EXPOSE 8080
CMD ["/app/filterid-ip-api"]
