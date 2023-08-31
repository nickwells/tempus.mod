package tempus_test

import (
	"testing"
	"time"

	"github.com/nickwells/tempus.mod/tempus"
	"github.com/nickwells/testhelper.mod/v2/testhelper"
)

func TestIsLeapYear(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		dt  time.Time
		exp bool
	}{
		{
			ID:  testhelper.MkID("regular leap year"),
			dt:  time.Date(2004, time.January, 1, 12, 0, 0, 0, time.UTC),
			exp: true,
		},
		{
			ID:  testhelper.MkID("regular non-leap year"),
			dt:  time.Date(2001, time.January, 1, 12, 0, 0, 0, time.UTC),
			exp: false,
		},
		{
			ID:  testhelper.MkID("4-century leap year"),
			dt:  time.Date(2000, time.January, 1, 12, 0, 0, 0, time.UTC),
			exp: true,
		},
		{
			ID:  testhelper.MkID("non-4-century non-leap year"),
			dt:  time.Date(1900, time.January, 1, 12, 0, 0, 0, time.UTC),
			exp: false,
		},
	}

	for _, tc := range testCases {
		if tempus.IsLeapYear(tc.dt) != tc.exp {
			t.Log(tc.IDStr())
			if tc.exp {
				t.Errorf("\t: %s should be a leap year but is not\n",
					tc.dt.Format("Mon 2006 Jan 02"))
			} else {
				t.Errorf("\t: %s should not be a leap year but is\n",
					tc.dt.Format("Mon 2006 Jan 02"))
			}
		}
	}
}

func TestDaysInMonth(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		dt  time.Time
		exp int
	}{
		{
			ID:  testhelper.MkID("January in leap year"),
			dt:  time.Date(2004, time.January, 1, 12, 0, 0, 0, time.UTC),
			exp: 31,
		},
		{
			ID:  testhelper.MkID("December in leap year"),
			dt:  time.Date(2004, time.December, 1, 12, 0, 0, 0, time.UTC),
			exp: 31,
		},
		{
			ID:  testhelper.MkID("February in leap year"),
			dt:  time.Date(2004, time.February, 1, 12, 0, 0, 0, time.UTC),
			exp: 29,
		},
		{
			ID:  testhelper.MkID("February in non-leap year"),
			dt:  time.Date(2005, time.February, 1, 12, 0, 0, 0, time.UTC),
			exp: 28,
		},
		{
			ID:  testhelper.MkID("January in non-leap year"),
			dt:  time.Date(2005, time.January, 1, 12, 0, 0, 0, time.UTC),
			exp: 31,
		},
		{
			ID:  testhelper.MkID("December in non-leap year"),
			dt:  time.Date(2005, time.December, 1, 12, 0, 0, 0, time.UTC),
			exp: 31,
		},
	}

	for _, tc := range testCases {
		if days := tempus.DaysInMonth(tc.dt); days != tc.exp {
			t.Log(tc.IDStr())
			t.Errorf("\t: %s should have %d days but has %d\n",
				tc.dt.Format("Jan in 2006"), tc.exp, days)
		}
	}
}

func TestDaysInYear(t *testing.T) {
	testCases := []struct {
		testhelper.ID
		dt  time.Time
		exp int
	}{
		{
			ID:  testhelper.MkID("leap year"),
			dt:  time.Date(2004, time.January, 1, 12, 0, 0, 0, time.UTC),
			exp: 366,
		},
		{
			ID:  testhelper.MkID("non-leap year"),
			dt:  time.Date(2005, time.January, 1, 12, 0, 0, 0, time.UTC),
			exp: 365,
		},
	}

	for _, tc := range testCases {
		if days := tempus.DaysInYear(tc.dt); days != tc.exp {
			t.Log(tc.IDStr())
			t.Errorf("\t: %s should have %d days but has %d\n",
				tc.dt.Format("2006"), tc.exp, days)
		}
	}
}
