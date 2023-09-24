package tempus

import (
	"testing"

	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestTimezoneNames(t *testing.T) {
	if HasTimezoneInfo() {
		names := TimezoneNames()
		testhelper.DiffInt[int](t, "TimezoneCount", "", len(names), tzCount)
	}
}
