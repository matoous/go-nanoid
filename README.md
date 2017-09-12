# Go Nanoid

This package is Go copy of [ai's](https://github.com/ai) [nanoid](https://github.com/ai/nanoid)!

**Safe.** It uses cryptographically strong random generator.

**Compact.** It uses more symbols than UUID (`A-Za-z0-9_~`)
and has the same number of unique options in just 22 symbols instead of 36.

## Install

Via go get tool

``` bash
$ go get github.com/matoous/go-nanoid
```

## Usage

Generate ID

``` go
id := gonanoid.Generate()
```
Change ID length
``` go
gonanoid.SetSize(32)
```
Change ID alphabet
``` go
gonanoid.SetAlphabt("abcde12345")
```

## Testing

``` bash
$ go test -c -i -o /tmp/TestGenerate_in_gonanoid_test_gogo gonanoid
```

## Notice

If you use Go Nanoid in your project, please let me know!

If you have any issues, just feel free and open it in this repository, thanks!

## Credits

- [ai](https://github.com/ai) - [nanoid](https://github.com/ai/nanoid)
- icza - his tutorial on [random strings in Go](https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
