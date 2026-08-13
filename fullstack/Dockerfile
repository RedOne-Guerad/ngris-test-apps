# --- build with the current patched Go toolchain (matches the platform's Go 1.25) ---
FROM golang:1.25-alpine AS build
WORKDIR /src
COPY go.mod ./
COPY main.go ./
RUN CGO_ENABLED=0 GOOS=linux go build -trimpath -ldflags="-s -w" -o /app .

# --- run on scratch: nothing to scan but the (patched) Go binary → no base-OS CVEs ---
FROM scratch
COPY --from=build /app /app
EXPOSE 8080
ENTRYPOINT ["/app"]
