package main

import (
	"context"
	"log"

	"fairwinds.com/anonymous-questions-app/pkg"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/glebarez/sqlite"
	"github.com/sirupsen/logrus"
	"github.com/uptrace/opentelemetry-go-extra/otelgorm"
	"go.opentelemetry.io/otel"
	"go.opentelemetry.io/otel/propagation"
	"gorm.io/gorm"

	"go.opentelemetry.io/contrib/instrumentation/github.com/gin-gonic/gin/otelgin"
	stdout "go.opentelemetry.io/otel/exporters/stdout/stdouttrace"
	sdktrace "go.opentelemetry.io/otel/sdk/trace"
)

func main() {
	db, err := gorm.Open(sqlite.Open("sqlite.db"), &gorm.Config{})
	if err != nil {
		logrus.Fatal("could not open db connection", err)
	}

	db.AutoMigrate(&pkg.Question{})

	err = db.Use(otelgorm.NewPlugin())
	if err != nil {
		logrus.Fatal("could not install otelgorm plugin", err)
	}

	h, err := pkg.NewHandler(db)
	if err != nil {
		logrus.Fatal("could not create handler", err)
	}

	r := gin.Default()
	r.Use(cors.Default())

	tp, err := initTracer()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		if err := tp.Shutdown(context.Background()); err != nil {
			log.Printf("Error shutting down tracer provider: %v", err)
		}
	}()

	r.Use(otelgin.Middleware("anonymous-questions-backend"))

	s, err := pkg.NewServer(r, h)
	if err != nil {
		logrus.Fatal("could not create server", err)
	}

	logrus.Fatal(s.Start())
}

func initTracer() (*sdktrace.TracerProvider, error) {
	exporter, err := stdout.New(stdout.WithPrettyPrint())
	if err != nil {
		return nil, err
	}
	tp := sdktrace.NewTracerProvider(
		sdktrace.WithSampler(sdktrace.AlwaysSample()),
		sdktrace.WithBatcher(exporter),
	)
	otel.SetTracerProvider(tp)
	otel.SetTextMapPropagator(propagation.NewCompositeTextMapPropagator(propagation.TraceContext{}, propagation.Baggage{}))
	return tp, nil
}
