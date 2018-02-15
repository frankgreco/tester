package options

import (
	"github.com/northwesternmutual/kanali/pkg/flags"
)

func init() {
	TesterOptions.Add(
		FlagServerSecurePort,
		FlagServerInsecurePort,
		FlagServerInsecureBindAddress,
		FlagServerSecureBindAddress,
		FlagServerTLSCertFile,
		FlagServerTLSKeyFile,
		FlagServerTLSCaFile,
	)
}

var (
	// FlagServerSecurePort sets the port that Kanali will listen on for incoming requests
	FlagServerSecurePort = flags.Flag{
		Long:  "server.secure_port",
		Short: "",
		Value: 0,
		Usage: "Sets the port that Kanali will listen on for incoming requests.",
	}
	// FlagServerInsecurePort sets the port that Kanali will listen on for incoming requests
	FlagServerInsecurePort = flags.Flag{
		Long:  "server.insecure_port",
		Short: "",
		Value: 0,
		Usage: "Sets the port that Kanali will listen on for incoming requests.",
	}
	// FlagServerInsecureBindAddress specifies the network address that Kanali will listen on for incoming requests
	FlagServerInsecureBindAddress = flags.Flag{
		Long:  "server.insecure_bind_address",
		Short: "",
		Value: "0.0.0.0",
		Usage: "Network address that Kanali will listen on for incoming requests.",
	}
	// FlagServerSecureBindAddress specifies the network address that Kanali will listen on for incoming requests
	FlagServerSecureBindAddress = flags.Flag{
		Long:  "server.secure_bind_address",
		Short: "",
		Value: "0.0.0.0",
		Usage: "Network address that Kanali will listen on for incoming requests.",
	}
	// FlagServerTLSCertFile specifies the path to x509 certificate for HTTPS servers.
	FlagServerTLSCertFile = flags.Flag{
		Long:  "server.tls.cert_file",
		Short: "c",
		Value: "",
		Usage: "Path to x509 certificate for HTTPS servers.",
	}
	// FlagServerTLSKeyFile pecifies the path to x509 private key matching --tls-cert-file
	FlagServerTLSKeyFile = flags.Flag{
		Long:  "server.tls.key_file",
		Short: "k",
		Value: "",
		Usage: "Path to x509 private key matching --tls.cert_file.",
	}
	// FlagServerTLSCaFile specifies the path to x509 certificate authority bundle for mutual TLS
	FlagServerTLSCaFile = flags.Flag{
		Long:  "server.tls.ca_file",
		Short: "",
		Value: "",
		Usage: "Path to x509 certificate authority bundle for mutual TLS.",
	}
)
