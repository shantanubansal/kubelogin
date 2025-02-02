//+build wireinject

// Package di provides dependency injection.
package di

import (
	"github.com/google/wire"
	"github.com/shantanubansal/kubelogin/pkg/cmd"
	"github.com/shantanubansal/kubelogin/pkg/credentialplugin/writer"
	"github.com/shantanubansal/kubelogin/pkg/infrastructure/browser"
	"github.com/shantanubansal/kubelogin/pkg/infrastructure/clock"
	"github.com/shantanubansal/kubelogin/pkg/infrastructure/logger"
	"github.com/shantanubansal/kubelogin/pkg/infrastructure/mutex"
	"github.com/shantanubansal/kubelogin/pkg/infrastructure/reader"
	"github.com/shantanubansal/kubelogin/pkg/infrastructure/stdio"
	kubeconfigLoader "github.com/shantanubansal/kubelogin/pkg/kubeconfig/loader"
	kubeconfigWriter "github.com/shantanubansal/kubelogin/pkg/kubeconfig/writer"
	"github.com/shantanubansal/kubelogin/pkg/oidc/client"
	"github.com/shantanubansal/kubelogin/pkg/tlsclientconfig/loader"
	"github.com/shantanubansal/kubelogin/pkg/tokencache/repository"
	"github.com/shantanubansal/kubelogin/pkg/usecases/authentication"
	"github.com/shantanubansal/kubelogin/pkg/usecases/credentialplugin"
	"github.com/shantanubansal/kubelogin/pkg/usecases/setup"
	"github.com/shantanubansal/kubelogin/pkg/usecases/standalone"
)

// NewCmd returns an instance of infrastructure.Cmd.
func NewCmd() cmd.Interface {
	wire.Build(
		NewCmdForHeadless,

		// dependencies for production
		clock.Set,
		stdio.Set,
		logger.Set,
		browser.Set,
	)
	return nil
}

// NewCmdForHeadless returns an instance of infrastructure.Cmd for headless testing.
func NewCmdForHeadless(clock.Interface, stdio.Stdin, stdio.Stdout, logger.Interface, browser.Interface) cmd.Interface {
	wire.Build(
		// use-cases
		authentication.Set,
		standalone.Set,
		credentialplugin.Set,
		setup.Set,

		// infrastructure
		cmd.Set,
		reader.Set,
		kubeconfigLoader.Set,
		kubeconfigWriter.Set,
		repository.Set,
		client.Set,
		loader.Set,
		writer.Set,
		mutex.Set,
	)
	return nil
}
