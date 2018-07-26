package cipd

import (
	"log"
	"os/user"
	"path/filepath"
	"runtime"

	"go.kendal.io/chow"
)

// TODO: Make this flexible enough to not rely on my local CIPD path.
func executable() string {
	usr, err := user.Current()
	if err != nil {
		log.Fatal(err)
	}

	path := usr.HomeDir + "/Downloads/a/depot_tools/cipd"
	if runtime.GOOS == "windows" {
		path += ".bat"
	}

	return filepath.FromSlash(path)
}

func Ensure(pkg string, ver string) chow.StepProvider {
	ensureFile := chow.Placeholder(pkg + " " + ver)
	return &chow.SelfProvider{
		Command: []string{
			executable(),
			"ensure",
			"-root",
			"//packages",
			"-ensure-file",
			ensureFile,
		},
		Outputs: []string{},
	}
}
