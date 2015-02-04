Franco
======

Detect the language of text.


## Installation

```
$ go get github.com/kapsteur/franco
```

## Usage

```go
res := franco.Detect("Votre temps est limité, ne le gâchez pas en menant une existence qui n’est pas la vôtre.")
// res == {Code:"fra" Count:1}
```

## TODO:

* Test, test and test again

## Derivation

Franco is a derivative work from [Franc](https://github.com/wooorm/franc) (Js, MIT).

## License

[MIT](LICENSE) © [Garry POUPIN](http://garry.io)