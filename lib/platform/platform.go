// Package platform provides Operating system and Architecture constants for the current
// platform, as they are used in CIPD.
package platform

import "runtime"

// TODO: System properties need to be mocked in tests.  Maybe these should be moved into
// go.kendal.io/chow.
const Arch = runtime.GOARCH
const OS = os
