# This is project to learn tutorial of Golang.

### [Get start with Hello World.](https://golang.org/doc/tutorial/create-module)

### [Call your code from another module](https://golang.org/doc/tutorial/call-module-code).

```go
// The command specifies that greetings should be
// replaced with ../greetings for the purpose of 
// locating the dependency

$ go mod edit -replace greetings =../greetings

// From the command prompt in the hello directory
// run the go mod tidy command to synchronize
// the hello module's dependencies, adding those 
// required by the code, but not yet tracked in the module.

$ go mod tidy
   ```

### [Add a test](https://golang.org/doc/tutorial/add-a-test)

```go
$ go test
// or
$ go test -v
```

### [Complie and install the application](https://golang.org/doc/tutorial/compile-install)

```go
// complies the packages
$ go build

// executable 
$ ./hello // with Linux or Mac
$ hello.exe // with Windows

// complies and install packges
$ go install

// excutable
$ hello
```