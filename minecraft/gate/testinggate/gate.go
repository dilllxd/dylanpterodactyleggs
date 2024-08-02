package main

import (
	"go.minekube.com/gate/cmd/gate"
	"go.minekube.com/gate/pkg/edition/java/proxy"
)

// It's a normal Go program, we only need
// to register our plugins and execute Gate.
func main() {
	// Here we register our plugins with the proxy.
	proxy.Plugins = append(proxy.Plugins) // We have some demo plugins in the plugins/ directory,
	// but you can also import your own plugins from other repositories.
	//
	// Checkout https://github.com/minekube/awesome for some inspiration.

	// Add more plugins as you like.
	// They will be initialized in the same order as appended.

	// Simply execute Gate as if it was a normal Go program.
	// Gate will take care of everything else for us,
	// such as config auto-reloading and flags like --debug.
	gate.Execute()
}
