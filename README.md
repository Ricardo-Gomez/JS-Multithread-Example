# JS-Multithread-Example
small Go server to show how JS implements multithread tasks in browser using web-workers

The Go server Help us setting the headers needed to use the web Workers in the browser to take advantage of multithreads
`"Cross-Origin-Opener-Policy: same-origin"`
`"Cross-Origin-Embedder-Policy: require-corp"`

## requirements

Go > 17
## install
```sh
$ go install
```

## run
```sh
$ go run server.go
```

## look around

> http://localhost:3000/gol
> Single thread Game of Life implementation.[^1]

> http://localhost:3000/gol-thread
> Multithread Game of Life implementation.[^1]

[^1]:Code source is from the book **Multithread Javascript** By Thomas Hunter and Bryan English.

