module github.com/slamy-solutions/openbp-go

go 1.21.4

require (
	github.com/redis/go-redis/v9 v9.0.2
	go.mongodb.org/mongo-driver v1.11.1
	go.opentelemetry.io/contrib/instrumentation/go.mongodb.org/mongo-driver/mongo/otelmongo v0.38.0
	go.opentelemetry.io/contrib/propagators/jaeger v1.13.0
	go.opentelemetry.io/otel v1.12.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace v1.12.0
	go.opentelemetry.io/otel/exporters/otlp/otlptrace/otlptracegrpc v1.12.0
	go.opentelemetry.io/otel/sdk v1.12.0
	go.opentelemetry.io/otel/trace v1.12.0
	google.golang.org/grpc v1.52.3
)

require (
	github.com/cenkalti/backoff/v4 v4.2.0 // indirect
	github.com/cespare/xxhash/v2 v2.2.0 // indirect
	github.com/dgryski/go-rendezvous v0.0.0-20200823014737-9f7001d12a5f // indirect
	github.com/go-logr/logr v1.2.3 // indirect
	github.com/go-logr/stdr v1.2.2 // indirect
	github.com/golang/snappy v0.0.1 // indirect
	github.com/grpc-ecosystem/grpc-gateway/v2 v2.7.0 // indirect
	github.com/klauspost/compress v1.15.15 // indirect
	github.com/montanaflynn/stats v0.0.0-20171201202039-1bf9dbcd8cbe // indirect
	github.com/nats-io/nats-server/v2 v2.9.12 // indirect
	github.com/nats-io/nkeys v0.3.0 // indirect
	github.com/nats-io/nuid v1.0.1 // indirect
	github.com/pkg/errors v0.9.1 // indirect
	github.com/xdg-go/pbkdf2 v1.0.0 // indirect
	github.com/xdg-go/scram v1.1.1 // indirect
	github.com/xdg-go/stringprep v1.0.3 // indirect
	github.com/youmark/pkcs8 v0.0.0-20181117223130-1be2e3e5546d // indirect
	go.opentelemetry.io/otel/exporters/otlp/internal/retry v1.12.0 // indirect
	go.opentelemetry.io/otel/metric v0.35.0 // indirect
	go.opentelemetry.io/proto/otlp v0.19.0 // indirect
	golang.org/x/crypto v0.5.0 // indirect
	golang.org/x/sync v0.0.0-20210220032951-036812b2e83c // indirect
)

require (
	github.com/golang/protobuf v1.5.2
	github.com/nats-io/nats.go v1.23.0
	github.com/sirupsen/logrus v1.9.0
	go.opentelemetry.io/contrib/instrumentation/google.golang.org/grpc/otelgrpc v0.38.0
	golang.org/x/net v0.5.0 // indirect
	golang.org/x/sys v0.4.0 // indirect
	golang.org/x/text v0.6.0 // indirect
	google.golang.org/genproto v0.0.0-20221118155620-16455021b5e6 // indirect
	google.golang.org/protobuf v1.28.1
)
