## versionsort

[![LICENSE](https://img.shields.io/badge/license-MIT-blue.svg)](LICENSE)
[![GoDoc](https://godoc.org/github.com/haya14busa/versionsort?status.svg)](https://godoc.org/github.com/haya14busa/versionsort)

### Usage

```go
import "github.com/haya14busa/go-versionsort"

func ExampleVersionSort() {
	strs := []string{
		"v1.1",
		"v1.10",
		"v1.11",
		"v1.9",
		"v1.8",
	}
	versionsort.Sort(strs, false)
	for _, s := range strs {
		fmt.Println(s)
	}
	// Output:
	// v1.1
	// v1.8
	// v1.9
	// v1.10
	// v1.11
}
```
