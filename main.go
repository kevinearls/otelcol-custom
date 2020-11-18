package main

import (
	"fmt"
	"log"

	"go.opentelemetry.io/collector/component"
	"go.opentelemetry.io/collector/service"
)

func main() {
	//factories, err := defaultcomponents.Components()
	factories, err := components()
	if err != nil {
		log.Fatalf("failed to build components: %v", err)
	}

	/*
	for k := range factories.Extensions {
		fmt.Printf("Extensions: %s\n", k)
	}
	for k := range factories.Exporters {
		fmt.Printf("Exporter: %s\n", k)
	}
	for k := range factories.Receivers {
		fmt.Printf("Receiver: %s\n", k)
	}*/

	fmt.Println("========== Processors ==========")
	for p := range factories.Processors {
		fmt.Printf("\t%s\n", p)
		pl := factories.Processors[p]
		fmt.Printf("\t\t is a %s\n", pl.Type())
	}
	fmt.Println("========== Exporter ==========")

	for e := range factories.Exporters {
		fmt.Printf("\t%s\n", e)
	}


	info := component.ApplicationStartInfo{
		ExeName:  "otelcol-custom",
		LongName: "Custom OpenTelemetry Collector distribution",
		Version:  "1.0.0",
	}

	app, err := service.New(service.Parameters{ApplicationStartInfo: info, Factories: factories})
	if err != nil {
		log.Fatal("failed to construct the application: %w", err)
	}

	err = app.Run()
	if err != nil {
		log.Fatal("application run finished with error: %w", err)
	}
}

