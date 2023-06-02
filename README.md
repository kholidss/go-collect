# go-collect

`go-collect` is a Go library inspired by Laravel Collection that provides convenient functions for working with collections, such as arrays, slices, maps, and more. It offers various utility functions to manipulate, transform, and analyze collections.

## Installation

To use Go-Collect in your Go project, you need to have Go installed and set up on your system. Then, you can install Go-Collect using the `go get` command:
```
go get github.com/kholidss/go-collect
```

## Usage

```
import "github.com/kholidss/go-collect"
```

## Example

```go
package main

import (
	"fmt"

	"github.com/kholidss/go-collect"
)

func main() {
	count, err := collect.Count("12")
	if err != nil {
		fmt.Printf("go-collect Count error: %v\n", err)
		return
	}
	fmt.Printf("count = %d\n", count)
	// Output:
	// 2

	count, err = collect.Count(map[string]interface{}{
		"name":      "John Doe",
		"age":       29,
		"isAllowed": true,
	})
	if err != nil {
		fmt.Printf("go-collect Count error: %v\n", err)
		return
	}
	fmt.Printf("count = %d\n", count)
	// Output:
	// 3

	countBy := collect.CountBy([]interface{}{
		"John", 1, 2, 1, false,
	})
	if err != nil {
		fmt.Printf("go-collect CountBy error: %v\n", err)
		return
	}
	fmt.Printf("countBy = %+v\n", countBy)
	// Output:
	// map[false:1 1:2 2:1 John:1]
}

```

## Contributing
Contributions to Go-Collect are welcome! If you find any issues or have suggestions for improvements, please open an issue or submit a pull request on this GitHub repository.

## License
Go-Collect is licensed under the [MIT License](https://choosealicense.com/licenses/mit/). Feel free to use, modify, and distribute this package according to the terms of the license.