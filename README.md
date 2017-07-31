Franco [![GoDoc](https://godoc.org/github.com/kapsteur/franco?status.png)](https://godoc.org/github.com/kapsteur/franco) [![Build Status](https://api.travis-ci.org/kapsteur/franco.svg)](https://travis-ci.org/kapsteur/franco)
======



Detect the language of text.


## Supported languages

Franco supports 175 “languages”. For a complete list, check out [languages list](https://github.com/kapsteur/franco/blob/master/languages.md)


## Installation

```
$ go get github.com/kapsteur/franco
```

## Usage

```go
res := franco.DetectOne("Votre temps est limité, ne le gâchez pas en menant une existence qui n’est pas la vôtre.")
// res == {Code:"fra" Count:1}

res := franco.Detect("Votre temps est limité, ne le gâchez pas en menant une existence qui n’est pas la vôtre.")
// res == [{Code:"fra" Count:1},{spa 0.7709821779068855},{cat 0.7656434011148622},{src 0.7274083379131664}...]
```

## Derivation

Franco is a derivative work from [Franc](https://github.com/wooorm/franc) (Js, MIT).
Franc is a derivative work from [guess-language][] (Python, LGPL),
[guesslanguage][] (C++, LGPL), and [Language::Guess][language-guess]
(Perl, GPL).  Their creators granted me the rights to distribute franc
under the MIT license: respectively, [Kent S. Johnson][grant-3],
[Jacob R. Rideout][grant-2], and [Maciej Ceglowski][grant-1].



## License

[MIT](LICENSE) © [Garry POUPIN](http://garry.io)
