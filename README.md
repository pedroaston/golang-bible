# golang-bible
Just documenting my golang journey. If you have any doubt about go, please create an issue. Enjoy!

## How to setup your golang environment
- Install golang into your machine (https://go.dev/doc/install)
- (Optional) My favorite IDE is VSCode. You can then install the most used golang extension to have the best possible golang development environment :smiley:
- In a empty directory create your go.mod file
```go
module project-name

go 1.19 // the go version you have

require(
    github.com/google/some-repo v0.1.3 // the repositories your project uses and the version being used
)
```
- create a *.go file with this structure in your directory
```go
package main

import "fmt"

func main() {
    fmt.Println("First step is done!")
}
```
- Start exploring the initiate directory