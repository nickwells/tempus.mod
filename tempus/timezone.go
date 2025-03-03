package tempus

import (
	"archive/zip"
	"os"
	"os/exec"
)

var zoneInfoFile string

func init() {
	goPath, err := exec.LookPath("go")
	if err != nil {
		return
	}

	out, err := exec.Command(goPath, "env", "GOROOT").Output()
	if err != nil {
		return
	}

	zoneInfoFile = string(out) + "/lib/time/zoneinfo.zip"
}

// Go1.21 has 597 timezone names
const tzCount = 597

// HasTimezoneInfo returns true if the timezone info file can be found
func HasTimezoneInfo() bool {
	fi, err := os.Stat(zoneInfoFile)
	if err != nil {
		return false
	}

	if fi.Size() == 0 {
		return false
	}

	return true
}

// TimezoneNames returns the list of names of valid timezones. It returns a
// nil slice if there are any problems finding the timezones. Problems
// include the timezones file not existing, being empty or the go executable
// not being in the PATH (so GOROOT cannot be determined)
func TimezoneNames() []string {
	if !HasTimezoneInfo() {
		return nil
	}

	r, err := zip.OpenReader(zoneInfoFile)
	if err != nil {
		return nil
	}

	zones := make([]string, 0, tzCount)
	for _, f := range r.File {
		zones = append(zones, f.Name)
	}

	return zones
}
