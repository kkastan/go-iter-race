# go-iter-race

Illustrates race conditions using go routines within an iteration with and without the Spackler library.

Try the following:

Doesn't use the Spackler library and correctly sets up a variable for the closure inside the iteration.
```
$ go run -race main.go
```

Doesn't use the Spackler library and does not use a separate value within the closure inside the iteration. This code will fail.
```
$ go run -race fail.go
```

Uses the Spackler library and correctly creates a new value within the iteration. This works.
```
$ go run -race spackler.go
```

Uses the Spackler library and does not use a separate value within the iteration. This does not work.
```
$ go run -race spackler_fail.go
```
