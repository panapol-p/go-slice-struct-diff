# go-slice-struct-listener

A Go library to find diff slice of struct (when feeding new data then output show id was added, id was updated or id was deleted)
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
Contributing Guidelines📝
Are we missing any of your favorite features, which you think you can add to it? For major changes, please open an issue first to discuss what you would like to change. You are welcome to contribute to this project.
To start contributing, follow the below guidelines:
1. Fork this repository.
2. Clone your forked copy of the project.
git clone https://github.com/<your_user_name>/go-slice-struct-diff.git
3. Navigate to the project directory 📁
cd e-commerce_redstore.github.io
4. Add a reference(remote) to the original repository.
git remote add upstream (https://github.com/panapol-p/go-slice-struct-diff.git)
5. Check the remotes for this repository.
git remote -v
6. Always take a pull from the upstream repository to your master branch to keep it at par with the main project(updated repository).
git pull upstream main
7. Create a new branch.
git checkout -b <your_branch_name>
8. Perform your desired changes to the code base.
9. Track your changes ✔️.
git add .
10. Commit your changes.
git commit -m "Relevant message"
11. Push the committed changes in your feature branch to your remote repo.
git push -u origin <your_branch_name>
12. To create a pull request, click on compare and pull requests.
13. Add an appropriate title and description to your pull request explaining your changes and efforts.
14. Click on Create Pull Request.
15. Woohoo!🥳 You have made a PR to the **go-slice-struct-diff**. Wait for your submission to be accepted and your PR to be merged.
You made it! 🎊

<a href="https://github.com/panapol-p/go-slice-struct-diff/graphs/contributors">
  <img src="https://contrib.rocks/image?repo=panapol-p/go-slice-struct-diff" />
</a>

<br>
<div align="center">
Show some ❤️ by starring this awesome repository!
</div>
