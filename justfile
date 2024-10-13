# Starts the infrastructure
start-infra:
  docker compose up -d

# Stops the infrastructure
stop-infra:
  docker compose down

# Starts the service
run-service:
  go run ./cmd

# Run k6 tests
run-tests:
  k6 run ./tests/load.js

# Quickstart
quickstart:
  just start-infra
  just run-service
