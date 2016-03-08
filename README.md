# iptc

IPTC reader - Go (golang) wrapper for libiptcdata

## Dependencies

Requires `libiptcdata`. On OS X using Homebrew:

```
brew install libiptcdata
```

## Usage

```go
package main

import (
        "fmt"
        "os"

        "github.com/zidizei/iptc"
)

func main() {
        data, _ := iptc.Open(os.Args[1])

        fmt.Printf("%v\n", data)
}
```

Output:

```
$ go run read.go testdata/sample.jpg
{testdata/sample.jpg map[Headline:iptc sample CopyrightNotice:Patrick 2016 ApplicationRecordVersion:4 Keywords:new tag sample first golang iptc sp€ciäl By-line:Patrick By-lineTitle:Gopher Contact:zidizei@gmail.com Writer-Editor:Patrick CodedCharacterSet:[1 b   2 5   4 7] EnvelopeRecordVersion:4 ObjectName:Sample 1]}
```
