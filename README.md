## Introduction to programming in Go

The purpose of this assignment is simply
to have you write a few lines of code in
Go and to exercise your git skills. Neither
the Go nor git tasks require any tech
ninja-ry.

### How to complete this assignment

The task in this assignment is to edit
the `main.go` file such that the program
takes one argument from
the command line (a number) and prints
your class nickname that many times.

For me, that would be something like
the following:

```
❯ go run main.go 7
bald-chicken
bald-chicken
bald-chicken
bald-chicken
bald-chicken
bald-chicken
bald-chicken
```

or

```
❯ go run main.go 3
bald-chicken
bald-chicken
bald-chicken
```

If the number provided is invalid (not an integer),
less than zero, greater than 10, or missing your program
should instead print "ERROR!".

```
❯ go run main.go -1
ERROR!
❯ go run main.go 11
ERROR!
❯ go run main.go 11.4
ERROR!
❯ go run main.go
ERROR!
```

To complete this assignment, you'll likely
want to look at the following resources:

- [Hello world in Go](https://go.dev/tour/welcome/1)
- [For loops in Go](https://go.dev/tour/flowcontrol/1)
- [Command line arguments in Go](https://gobyexample.com/command-line-arguments)
- [the Go `strconv.Atoi` function](https://pkg.go.dev/strconv#Atoi)

Right now, when you run `go run main.go`, it will
print "नमस्ते, 世界.", which is "namaste shijie": "hello"
in Hindi and "world" in Chinese. I provided that code
only as a placeholder.

The assignment should take you less than 30 minutes
and it should only require about 10 lines of code.
If you write more than 30 lines of code you're doing
something really weird and you should see a TA 😜.
