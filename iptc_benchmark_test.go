package iptc

import "testing"

func BenchmarkOpen(b *testing.B) {
        var file = "testdata/sample.jpg"
        for i := 0; i < b.N; i++ {
                Open(file)
        }
}

func BenchmarkLoad(b *testing.B) {
        type Meta struct {
                Keywords   []string
                ObjectName string
                ApplicationRecordVersion int
        }

        var file = "testdata/sample.jpg"
        for i := 0; i < b.N; i++ {
                actual := Meta{}
                Load(file, &actual)
        }
}
