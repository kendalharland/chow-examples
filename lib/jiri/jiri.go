// Package jiri provides support for working with the Jiri command.
//
// TODO: Declare outputs
package jiri

import (
	"fmt"

	"go.kendal.io/chow"
	"go.kendal.io/chow-examples/lib/cipd"
	"go.kendal.io/chow-examples/lib/platform"
)

const (
	// The location where Jiri files are installed.
	installRoot = "//packages/jiri"

	// Path to the Jiri executable.
	executable = "//packages/jiri/jiri"
)

// Installs Jiri from CIPD.
func EnsureJiri(version string) chow.Step {
	pkg := fmt.Sprintf("fuchsia/tools/jiri/%s-%s", platform.OS, platform.Arch)
	return cipd.Ensure(installRoot, map[string]string{pkg: version})
}

// Initializes a Jiri root in the current directory.
func Init(dir string) chow.Step {
	return chow.Step{
		Command: []string{
			executable,
			"init",
			"-v",
			"-time",
			"-analytics-opt=false",
			"-rewrite-sso-to-https=true",
			"-cache", "//cache/git",
			"-shared",
		},
	}
}

// Imports a project into the specified manifest file.
//
// TODO: overwrite
func Import(manifest, remote, name, revision string) chow.Step {
	cmd := []string{executable, "import", "-v", "-time"}

	if name != "" {
		cmd = append(cmd, "-name", name)
	}
	if revision != "" {
		cmd = append(cmd, "-revision", revision)
	}

	cmd = append(cmd, manifest, remote)
	return chow.Step{Command: cmd}
}

// Updates the local checkout according to the Jiri manifest.
//
// TODO: rebase_tracked, local_manifest, snapshot.
func Update(gc bool, runHooks bool, attempts int) chow.Step {
	if attempts == 0 {
		attempts = 3
	}

	cmd := []string{
		executable,
		"update",
		"-v",
		"-time",
		"-autoupdate=false",
		"-attempts", fmt.Sprintf("%d", attempts),
	}
	if gc {
		cmd = append(cmd, "-gc=true")
	}
	if !runHooks {
		cmd = append(cmd, "-run-hooks=false")
	}

	return chow.Step{Command: cmd}
}

func RunHooks(attempts int) chow.Step {
	if attempts == 0 {
		attempts = 3
	}

	return chow.Step{
		Command: []string{
			executable,
			"run-hooks",
			"-attempts", fmt.Sprintf("%d", attempts),
		},
	}
}

// func SourceManifest() chow.Step {}
// func Snapshot() chow.Step {}
