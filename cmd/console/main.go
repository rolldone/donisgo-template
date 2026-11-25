package main

import (
	"syscall"

	"github.com/rolldone/donisgo/console"

	// Register the starter's own console commands (plugin new) onto the
	// framework root command.
	_ "github.com/rolldone/donisgo-template/console"
)

func main() {
	// Ensure console commands create group-writable files by default
	syscall.Umask(0o002)
	// To register additional plugins and their console commands:
	// import "github.com/rolldone/donisgo/plugins/myplugin"
	// console.RegisterAdditionalPlugins([]plugins.Plugin{myplugin.New()})
	console.Execute()
}
