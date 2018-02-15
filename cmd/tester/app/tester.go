package app

import (
	"context"

	"github.com/northwesternmutual/kanali/pkg/chain"
	"github.com/northwesternmutual/kanali/pkg/server"
	"github.com/oklog/run"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"github.com/spf13/viper"

	"github.com/frankgreco/tester/cmd/tester/app/options"
	"github.com/frankgreco/tester/pkg/log"
	"github.com/frankgreco/tester/pkg/middleware"
)

// Run is this application's main wrapper
func Run(ctx context.Context) error {
	ctx, cancel := context.WithCancel(ctx)
	logger := log.WithContext(nil)

	tester := server.Prepare(&server.Options{
		Name:         "tester",
		InsecureAddr: viper.GetString(options.FlagServerInsecureBindAddress.GetLong()),
		SecureAddr:   viper.GetString(options.FlagServerSecureBindAddress.GetLong()),
		InsecurePort: viper.GetInt(options.FlagServerInsecurePort.GetLong()),
		SecurePort:   viper.GetInt(options.FlagServerSecurePort.GetLong()),
		TLSKey:       viper.GetString(options.FlagServerTLSKeyFile.GetLong()),
		TLSCert:      viper.GetString(options.FlagServerTLSCertFile.GetLong()),
		TLSCa:        viper.GetString(options.FlagServerTLSCaFile.GetLong()),
		Handler: chain.New().Add(
			middleware.Correlation,
			middleware.Metrics,
		).Link(middleware.Dump),
		Logger: logger.Sugar(),
	})

	metrics := server.Prepare(&server.Options{
		Name:         "prometheus",
		InsecureAddr: viper.GetString(options.FlagPrometheusInsecureBindAddress.GetLong()),
		InsecurePort: viper.GetInt(options.FlagPrometheusInsecurePort.GetLong()),
		Handler:      promhttp.Handler(),
		Logger:       logger.Sugar(),
	})

	var g run.Group

	g.Add(func() error {
		<-ctx.Done()
		return nil
	}, func(error) {
		cancel()
	})

	g.Add(func() error {
		return logError(metrics.Run())
	}, func(error) {
		logError(metrics.Close())
	})

	g.Add(func() error {
		return logError(tester.Run())
	}, func(error) {
		logError(tester.Close())
	})

	return g.Run()
}

func logError(err error) error {
	if err != nil {
		log.WithContext(nil).Error(err.Error())
	}
	return err
}
