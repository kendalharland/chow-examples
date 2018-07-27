// TODO: Declare outputs
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

// Install the given packges from CIPD.Ensure.
//
// packageVersions maps package paths to the versions that should be installed.
// root is the directory to install files to.
//
// Example:  Ensure(map[string]string{"infra/tools/gsutil": "latest"})
func Ensure(root string, packageVersions map[string]string) chow.Step {
	// Create ensure file.
	contents := new(bytes.Buffer)
	for pkg, ver := range packageVersions {
		fmt.Fprintln(contents, pkg+" "+ver)
	}
	ensureFile := chow.Placeholder(contents.String())
	jsonOutput := chow.Placeholder("")

	// Install packages.
	return chow.Step{
		Command: []string{
			executable(),
			"ensure",
			"-root", root,
			"-ensure-file", ensureFile,
			"-json-output", jsonOutput,
		},
	}
}
