package iptc

import "fmt"

func ExampleOpen() {
        data, err := Open("testdata/sample.jpg")
        if err != nil {
                panic(err)
        }

        str, err := data.StringTag("ObjectName")
        if err != nil {
                panic(err)
        }

        fmt.Printf("%s\n", str)
        // Output:
        // Sample 1
}

func ExampleLoad() {
        // Check http://www.sno.phy.queensu.ca/~phil/exiftool/TagNames/IPTC.html
        // for a list of available IPTC tags.
        type Data struct {
                ObjectName                string
                Keywords                  []string
                ApplicationRecordVersion  int
        }

        info := Data{}
        err  := Load("testdata/sample.jpg", &info)
        if err != nil {
                panic(err)
        }

        fmt.Printf("%#v\n", info)
        // Output:
        // iptc.Data{ObjectName:"Sample 1", Keywords:[]string{"new tag", "sample", "first", "golang", "iptc", "sp€ciäl"}, ApplicationRecordVersion:4}
}
