package iptc

import (
        "testing"
        "strings"
)

func TestLoadIptcData(t *testing.T) {
        type Meta struct {
                Keywords   []string
                ObjectName string
        }

        var tests = []struct {
                file      string
                expected  Meta
        }{
                {"testdata/sample.jpg",  Meta{[]string{"new tag", "sample", "first", "golang", "iptc", "sp€ciäl"}, "Sample 1"}},
                {"testdata/sample2.jpg", Meta{[]string{"hällo", "new tag", "sample", "golang", "iptc", "second"}, "Sample 2"}},
                {"testdata/sample3.jpg", Meta{[]string{"golang", "sample", "føø", "new tag", "third", "sp€ciäl", "iptc"}, "Sample 3"}},
                {"testdata/sample4.jpg", Meta{[]string{}, "Sample 4"}},
        }

        for _, c := range tests {
                actual := Meta{}
                err    := Load(string(c.file), &actual)

                if err != nil {
                        t.Errorf("Load(%q) failed: %s", c.file, err)
                }

                if actual.ObjectName != c.expected.ObjectName {
                        t.Errorf("expected [\"%s\"].ObjectName == '%s', got '%s'", c.file, c.expected.ObjectName, actual)
                }

                if len(actual.Keywords) != len(c.expected.Keywords) {
                        t.Errorf("expected [\"%s\"].Keywords to contain '%d' elements, got '%d'", c.file, len(c.expected.Keywords), len(actual.Keywords))
                }

        }
}

func TestParseIptcData(t *testing.T) {
        var tests = []struct {
                file, keyword, title string
        }{
                {"testdata/sample.jpg",  "new tag,sample,first,golang,iptc,sp€ciäl", "Sample 1"},
                {"testdata/sample2.jpg", "hällo,new tag,sample,golang,iptc,second", "Sample 2"},
                {"testdata/sample3.jpg", "golang,sample,føø,new tag,third,sp€ciäl,iptc", "Sample 3"},
                {"testdata/sample4.jpg", "", "Sample 4"},
        }

        for _, c := range tests {
                meta, err := Open(string(c.file))
                if err != nil {
                        t.Errorf("Open(%q) failed: %s", c.file, err)
                }

                keywords, err := meta.ListTag("Keywords")
                if err == nil {
                        actual := strings.Join(keywords, ",")

                        if actual != c.keyword {
                                t.Errorf("meta.ListTag(\"keywords\") == %q, expected %q", actual, c.keyword)
                        }
                } else {
                        // `sample4.jpg` does not have any Keywords, which means the `keywords` field is nil
                        if c.file != "testdata/sample4.jpg" && keywords != nil {
                                t.Errorf("meta.ListTag(\"keywords\") == %q, expected <nil>", keywords)
                        }
                }

                actual, err := meta.StringTag("ObjectName")
                if err != nil {
                        t.Errorf("meta.StringTag(\"ObjectName\") failed: %s", err)
                }

                if actual != c.title {
                        t.Errorf("meta.StringTag(\"ObjectName\") == %q, expected %q", actual, c.title)
                }
        }
}
