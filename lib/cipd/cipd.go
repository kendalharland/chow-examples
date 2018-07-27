package cipd

import (
	"bytes"
	"fmt"
	"path/filepath"
	"runtime"

	"go.kendal.io/chow"
)

// Expects Chrome depot tools' cipd[.bat] to be on the current PATH.
func executable() string {
	path := "cipd"
	if runtime.GOOS == "windows" {
		path += ".bat"
	}

	return filepath.FromSlash(path)
}

func Ensure(packageVersions map[string]string) chow.StepProvider {
	// Create ensure file.
	contents := new(bytes.Buffer)
	for pkg, ver := range packageVersions {
		fmt.Fprintln(contents, pkg+" "+ver)
	}
	ensureFile := chow.Placeholder(contents.String())

	// Install packages.
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
