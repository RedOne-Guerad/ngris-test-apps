# --- build the backend (no external deps → offline) ---
FROM golang:1.23-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -o /app .

# --- run: distroless, non-root ---
FROM gcr.io/distroless/static-debian12:nonroot
COPY --from=build /app /app
EXPOSE 8080
USER nonroot:nonroot
ENTRYPOINT ["/app"]
