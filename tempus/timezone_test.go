package tempus

import (
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestTimezoneNames(t *testing.T) {
	if HasTimezoneInfo() {
		names := TimezoneNames()
		testhelper.DiffInt(t, "TimezoneCount", "", len(names), tzCount)
	}
}
