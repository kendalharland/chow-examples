package main

import (
	"testing"

	"go.kendal.io/chow"
)

func TestMain(t *testing.T) {
	cfg := chow.TestConfig{Runnable: RunSteps}

	t.Run("default", func(t *testing.T) {
		cfg.Run(t, chow.TestCase{})
	})
}
