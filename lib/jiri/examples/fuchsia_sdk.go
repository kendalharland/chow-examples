package main

import (
	"go.kendal.io/chow"
	"go.kendal.io/chow-examples/lib/jiri"
)

func main() {
	chow.Main(RunSteps)
}

func RunSteps(r chow.Runner) {
	r.Run("ensure jiri", jiri.EnsureJiri("stable"))
	r.Run("jiri init", jiri.Init("//CWD/"))
	r.Run("jiri import topaz", jiri.Import(
		"manifest/topaz",
		"https://fuchsia.googlesource.com/topaz",
		"topaz",
		"c22471f4e3f842ae18dd9adec82ed9eb78ed1127",
	))
	r.Run("jiri update", jiri.Update(false, false, 3))
	r.Run("jiri run-hooks", jiri.RunHooks(3))
}
