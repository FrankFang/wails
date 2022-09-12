//go:build darwin && !bindings
// +build darwin,!bindings

package app

import (
	"github.com/frankfang/wails/v2/internal/logger"
	"github.com/frankfang/wails/v2/pkg/options"
)

func PreflightChecks(options *options.App, logger *logger.Logger) error {

	_ = options

	return nil
}
