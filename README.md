# This is project to learn tutorial of Golang.

1. [Get start with Hello World.](https://golang.org/doc/tutorial/create-module)
2. [Call your code from another module](https://golang.org/doc/tutorial/call-module-code).

```
// The command specifies that greetings should be
// replaced with ../greetings for the purpose of 
// locating the dependency

$ go mod edit -replace greetings=../greetings

// From the command prompt in the hello directory
// run the go mod tidy command to synchronize
// the hello module's dependencies, adding those 
// required by the code, but not yet tracked in the module.

$ go mod tidy
   ```