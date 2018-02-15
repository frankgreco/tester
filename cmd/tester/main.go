package main

import (
	"math/rand"
	"os"
	"runtime"
	"time"

	"github.com/northwesternmutual/kanali/pkg/flags"
	"github.com/northwesternmutual/kanali/pkg/log"
	"github.com/northwesternmutual/kanali/pkg/server"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"

	"github.com/frankgreco/tester/cmd/tester/app"
	"github.com/frankgreco/tester/cmd/tester/app/options"
)

var cmd = &cobra.Command{
	Use:   "tester",
	Short: "tester",
	Long:  "tester",
	Run:   run,
}

func run(cmd *cobra.Command, args []string) {
	if err := flags.InitViper("tester"); err != nil {
		log.WithContext(nil).Fatal(err.Error())
		os.Exit(1)
	}

	ctx := server.SetupSignalHandler()
	log.SetLevel(viper.GetString(options.FlagProcessLogLevel.GetLong()))
	if err := app.Run(ctx); err != nil {
		log.WithContext(nil).Error(err.Error())
		os.Exit(1)
	}
}

func main() {
	rand.Seed(time.Now().UTC().UnixNano())
	runtime.GOMAXPROCS(runtime.NumCPU())

	if err := options.TesterOptions.AddAll(cmd); err != nil {
		log.WithContext(nil).Fatal(err.Error())
		os.Exit(1)
	}

	if err := cmd.Execute(); err != nil {
		os.Exit(1)
	}
}
