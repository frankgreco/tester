package options

import (
	"github.com/northwesternmutual/kanali/pkg/flags"
)

func init() {
	TesterOptions.Add(
		FlagPrometheusInsecurePort,
		FlagPrometheusInsecureBindAddress,
	)
}

var (
	// FlagPrometheusInsecurePort sets the port that the Prometheus server will listen on for incoming requests
	FlagPrometheusInsecurePort = flags.Flag{
		Long:  "prometheus.insecure_port",
		Short: "",
		Value: 8000,
		Usage: "Sets the port that the Prometheus server will listen on for incoming requests.",
	}
	// FlagPrometheusInsecureBindAddress specifies the network address that the Prometheus server will listen on for incoming requests
	FlagPrometheusInsecureBindAddress = flags.Flag{
		Long:  "prometheus.insecure_bind_address",
		Short: "",
		Value: "0.0.0.0",
		Usage: "Network address that the Prometheus server will listen on for incoming requests.",
	}
)
