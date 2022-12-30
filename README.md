# Intro-To-Go-Training
GoLang Training December 2022

# Homework

The homework consists of Go source files including unimplemented functions.
Your job is to implement these functions based on their function comments and get the unit tests passing.
You will find `// TODO` comments in all the places you need to add code.

## Install Go

To do the homework you will need to install Go 1.18+ on your local machine.
You can do this with your package manager of choice.
Or you can download an installer (or tar archive) by following the instructions on the Go website: https://go.dev/doc/install

## How to verify your answers locally

To check your answers to one directory of problems, run all the unit tests in that directory.
The unit tests can be run with `go test`.

For example, you can run the tests for `001-slices` with this command (from the root dir of this Git repo):

```
go test 001-slices/*
```

### Windows Note:

If you are running the tests on Windows, the above command may not work.
If that is the case, please try running the above command (still from
the root dir of this Git repo) using the absolute path to the test folder.
For example:

```
go test C:\\Users\rakuten.taro\workspace\introduction-to-go-training\001-slices\
```

### 009-channels Note:

For the tests in `009-channels` you may want to use the `-timeout` flag.
Because these tests are concurrent, if there is a bug in your code it may lead to a deadlock.
Here's an example with a timeout of 10 seconds (which should be reasonable).

```
go test -timeout 10s 009-channels/*
```

## When you're done...

When you finish the homework, push your branch open a pull request to the `master` branch.
(Don't forget to `go fmt` your code first.)
The instructor will check your code and leave comments - or let you know if the code looks good.
If you need help or can't complete any problems, open a pull request and let the instructor know.
The instructor will check your code and leave comments to help you solve your problem.

## Useful commands

Here are some useful go CLI commands:

```
# Format source file
go fmt ${PATH_TO_SOURCE_FILE}

# Format all source files (run this in the module's root directory)
go fmt ./...

# Verify that a source file compiles
go build ${PATH_TO_SOURCE_FILE}

# Run unit tests
go test ${PATH_TO_TEST_FILE_AND_TESTED_SOURCE_FILE}
```
