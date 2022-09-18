package wininterface

import (
	"net"
	"sort"
	"testing"
)

func TestWinInterface(t *testing.T) {
	var expected []string
	var tests []string

	nifs, err := net.Interfaces()
	if err != nil {
		t.Error(err)
	}

	cmd := GetMac()
	macs := cmd.Parse()

	for _, nif := range nifs {
		for _, m := range macs {
			if nif.Name == m.ConnectionName {
				expected = append(expected, nif.Name)
			}
		}
	}

	for _, m := range macs {
		tests = append(tests, m.ConnectionName)
	}

	sort.Strings(expected)
	sort.Strings(tests)

	for i, e := range expected {
		if e != tests[i] {
			t.Errorf("expected %s but got %s", e, tests[i])
		}
	}

}
