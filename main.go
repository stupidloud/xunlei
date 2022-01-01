package main

import (
	"context"
	"os/signal"
	"syscall"

	"github.com/cnk3x/flags"
)

func main() {
	ctx, cancel := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM, syscall.SIGKILL)
	defer cancel()

	app := flags.New()
	app.AddCommand("run", &XunleiDaemon{})
	app.AddCommand("install", &Install{})
	app.AddCommand("uninstall", &Uninstall{})
	app.Run(ctx)
}
