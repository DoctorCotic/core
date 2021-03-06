package main

import (
	"context"
	"fmt"

	"github.com/noxiouz/zapctx/ctxlog"
	"github.com/sonm-io/core/cmd"
	"github.com/sonm-io/core/insonmnia/logging"
	"github.com/sonm-io/core/optimus"
)

var (
	configFlag  string
	versionFlag bool
	appVersion  string
)

func main() {
	cmd.NewCmd("optimus", appVersion, &configFlag, &versionFlag, run).Execute()
}

func run() error {
	cfg, err := optimus.LoadConfig(configFlag)
	if err != nil {
		return fmt.Errorf("failed to parse config: %v", err)
	}

	log := logging.BuildLogger(*cfg.Logging.Level)

	ctx := ctxlog.WithLogger(context.Background(), log)
	bot, err := optimus.NewOptimus(*cfg, log)
	if err != nil {
		return fmt.Errorf("failed to create Optimus: %v", err)
	}

	return bot.Run(ctx)
}
