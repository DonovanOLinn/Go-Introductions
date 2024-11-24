Every program in Go starts with a package declaration.

```go
package main
```
And this pints to the function that will be executed first.

```go
func main() {
	fmt.Println("Hello, world!")
}
```

Go is a modular language so we can separate the program into multiple files and compile them together. But the main function must be in the package main.

Then, we need to also import any other packages that we need.

```go
import "fmt"
```

When we call this import statement, we are importing the fmt package, which contains the Println function.

That is how this code works: 

```go
package main

import "fmt"

func main() {
	fmt.Println("Hello, world!")
}
````    

To run this code, we need to create a go.mod file.

```bash
go mod init hello
```

A GO file must import packages that are used in the file. It will not run if you have something that is imported but not used.
