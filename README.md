# go-slice-struct-diff

A Go library to find diff slice of struct (when feed new data then output show id was added , id was updated or id was deleted)
<hr>



## Usage <a id="usage"></a>
```go
package main

import (
	"fmt"

	diff "github.com/panapol-p/go-slice-struct-diff"
)

type FeedData struct {
	ID    string `diff:"id"`
	Name  string
	Score float32
}

func main() {
	fs := []FeedData{
		{ID: "1", Name: "Bob", Score: 98.50},
		{ID: "2", Name: "Joe", Score: 92.50},
	}

	d := diff.NewDiff[FeedData]()

	// set callback func if you need
	f := func(e []diff.Events[FeedData]) {
		fmt.Println("[callback func]", "receive new event!!", e)
	}
	d.SetCallback(f)

	events := d.AddNewValue(fs)
	fmt.Println(events) // [{1 added {1 Bob 98.5}} {2 added {2 Joe 92.5}}]

	fs = []FeedData{
		{ID: "1", Name: "Bob", Score: 96.50},
		{ID: "2", Name: "Joe", Score: 92.50},
		{ID: "3", Name: "Micky", Score: 89.70},
	}
	events = d.AddNewValue(fs)
	fmt.Println(events) // [{1 updated {1 Bob 96.5}} {3 added {3 Micky 89.7}}]

	fs = []FeedData{
		{ID: "1", Name: "Bob", Score: 96.50},
	}
	events = d.AddNewValue(fs)
	fmt.Println(events) // [{2 deleted {  0}} {3 deleted {  0}}]
}
```

## License <a id="license"></a>
Distributed under the MIT License. See [license](LICENSE) for more information.

## Contributing <a id="contributing"></a>
Contributions are welcome! Feel free to check our [open issues](https://github.com/panapol-p/go-slice-struct-listener/issues).

<a href="https://github.com/panapol-p/go-slice-struct-diff/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=panapol-p/go-slice-struct-diff" />
</a>

<br>
<div align="center">
Show some ❤️ by starring this awesome repository!
</div>
