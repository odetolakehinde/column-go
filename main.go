// Package main execution point for this integration
package main

import (
	"fmt"
	"os"

	"github.com/rs/zerolog"

	apiCalls "github.com/ovalfi/integrations-playground/column/api"
	"github.com/ovalfi/integrations-playground/column/model"
)

func main() {
	fmt.Println("Running Column SDK")
	_ = os.Setenv("TZ", "Africa/Lagos")
	logger := zerolog.New(os.Stderr).With().Timestamp().Logger()
	logger.Info().Msg("app starting")
	defer logger.Info().Msg("stopped")
	api := apiCalls.New(&logger, model.BaseURL, model.APIKey, model.ProxyIPAddressAndPort)
	api.RunInSandboxMode() // to ensure it is running in sandbox mode

	//ctx := context.Background()

}
