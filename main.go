package main

import (
	"context"
	"os"

	"github.com/shantanubansal/kubelogin/pkg/di"
)

var version = "HEAD"

func main() {
	os.Exit(di.NewCmd().Run(context.Background(), os.Args, version))
}
