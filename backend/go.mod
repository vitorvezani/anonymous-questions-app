module fairwinds.com/anonymous-questions-app

go 1.16

require (
	github.com/gin-contrib/cors v1.4.0
	github.com/gin-gonic/gin v1.8.2
	github.com/glebarez/sqlite v1.6.0
	github.com/sirupsen/logrus v1.9.0
	github.com/stretchr/testify v1.8.1
	github.com/uptrace/opentelemetry-go-extra/otelgorm v0.1.21
	go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin v0.39.0
	go.opentelemetry.io/otel v1.13.0
	go.opentelemetry.io/otel/exporters/jaeger v1.13.0 // indirect
	go.opentelemetry.io/otel/exporters/stdout/stdouttrace v1.13.0
	go.opentelemetry.io/otel/sdk v1.13.0
	gorm.io/gorm v1.24.5
)
