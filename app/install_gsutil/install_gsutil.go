package main

import (
	"go.kendal.io/chow"
	"go.kendal.io/chow-examples/lib/cipd"
)

func main() {
	chow.Main(RunSteps)
}

func RunSteps(r chow.Runner) {
	r.Run("install_gsutil", cipd.Ensure("infra/tools/gsutil", "latest"))
}
