package main

import (
	"os"

	"github.com/prometheus/client_golang/prometheus"

	"github.com/jessevdk/go-flags"
)

func main() {
	options := struct {
		Verbose        []bool            `short:"v" environment:"VERBOSE" long:"verbose" description:"Verbose output"`
		MetricsPath    string            `long:"metrics.path" environment:"METRICS_PATH" default:"/metrics" description:"Endpoint for provide metrics"`
		Listen         string            `long:"listen" environment:"LISTEN" default:":8080" description:"Interface for listening"`
		StatusFile     string            `long:"status-file" environment:"STATUS_FILE" default:"./status" description:"Status file for state"`
		ConstantLabels map[string]string `long:"labels" environment:"LABELS" description:"Constant labels map" default:"enviroment:dev"`
	}{}

	parser := flags.NewParser(&options, flags.Default)
	if _, err := parser.Parse(); err != nil {
		if flagsErr, ok := err.(*flags.Error); ok && flagsErr.Type == flags.ErrHelp {
			os.Exit(0)
		} else {
			os.Exit(1)
		}
	}
	prometheus.Collector
}
