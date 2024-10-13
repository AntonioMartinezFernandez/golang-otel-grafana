# golang-otel-grafana

Instrumentalized Golang service with OpenTelemetry and Grafana

## Resources

- [OpenTelemetry](https://opentelemetry.io/)
- [Grafana](https://grafana.com/)
- [OTEL collector example with Jaeger and Prometheus](https://github.com/open-telemetry/opentelemetry-go-contrib/tree/main/examples/otel-collector)
- [OpenTelemetry and Grafana in a Docker image](https://grafana.com/blog/2024/03/13/an-opentelemetry-backend-in-a-docker-image-introducing-grafana/otel-lgtm/)
- [Send logs with OTEL SDK](https://docs.honeycomb.io/send-data/logs/opentelemetry/sdk/go/)

## Get started

1. Install [Docker](https://docs.docker.com/get-docker/)
1. Install [Just](https://github.com/casey/just)
1. Install [k6](https://grafana.com/docs/k6/latest/set-up/install-k6/)
1. Run `just start-infra` to start the infrastructure
1. Run `just run-service` to start the example service
1. Run `just run-tests` to run the load tests and generate logs, metrics and traces
1. Use the Grafana UI to view the observability data (http://localhost:3000)
