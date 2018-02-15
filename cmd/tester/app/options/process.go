package options

import (
	"github.com/northwesternmutual/kanali/pkg/flags"
)

func init() {
	TesterOptions.Add(
		FlagProcessLogLevel,
	)
}

var (
	// FlagProcessLogLevel sets the logging level. Choose between 'debug', 'info', 'warn', 'error', 'fatal'
	FlagProcessLogLevel = flags.Flag{
		Long:  "process.log_level",
		Short: "l",
		Value: "info",
		Usage: "Sets the logging level. Choose between 'debug', 'info', 'warn', 'error', 'fatal'.",
	}
)
