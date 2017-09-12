# Go-Nanoid

This package is Go copy of [ai's](https://github.com/ai) [nanoid](https://github.com/ai/nanoid)!

## Install

Via go get tool

``` bash
$ go get github.com/matoous/go-nanoid
```

## Usage

``` go
# construct client
id := gonanoid.Generate()
# change alphabet
gonanoid.SetAlphabt("abcde12345")
# change id size
gonanoid.SetSize(16)
```

## Testing

``` bash
$ go test -c -i -o /tmp/TestGenerate_in_gonanoid_test_gogo gonanoid
```

## Notice

If you have any issues, just feel free and open it in this repository, thx!

## Credits

- [ai](https://github.com/ai) - [nanoid](https://github.com/ai/nanoid)
- icza - his tutorial on [random strings in Go](https://stackoverflow.com/questions/22892120/how-to-generate-a-random-string-of-a-fixed-length-in-golang)

## License

The MIT License (MIT). Please see [License File](LICENSE.md) for more information.
