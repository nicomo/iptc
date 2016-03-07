package iptc

import (
        "testing"
        "strings"
)

func TestParseIptcKeywords(t *testing.T) {
        var tests = []struct {
                file, expected string
        }{
                {"testdata/sample.jpg", "blackandwhite,people,street,urban,zurich"},
        }

        for _, c := range tests {
                meta, err := Open(string(c.file))
                actual    := strings.Join(meta["IPTCApplication"].Data["Keywords"], ",")

                if err != nil {
                        t.Errorf("Open(%q) failed: %s", c.file, err)
                }

                if actual != c.expected {
                        t.Errorf("Open(%q)[2][25] == %q, expected %q", c.file, actual, c.expected)
                }
        }
}
