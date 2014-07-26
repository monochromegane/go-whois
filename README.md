# go-whois

WHOIS client in GO.

## Features

- Print WHOIS response in CLI.
- Check domain exist in go package.

## Installation

```sh
$ go get github.com/monochromegane/go-whois/...
```

## Usage

### CLI

```sh
$ go-whois DOMAIN_NAME
```

### Package

```go
import whois

res, err := whois.Query(os.Args[1])
if err != nil {
	fmt.Println(err)
}
res.Raw()   // get original WHOIS response.
res.Exist() // check domain exist.
```

## Tasks

- Pasrse more WHOIS attributes.
- Support more WHOIS servers.

## Code status

- [![Build Status](https://travis-ci.org/monochromegane/go-whois.svg?branch=master)](https://travis-ci.org/monochromegane/go-whois)

## Contributing

1. Fork it
2. Create your feature branch (`git checkout -b my-new-feature`)
3. Commit your changes (`git commit -am 'Add some feature'`)
4. Push to the branch (`git push origin my-new-feature`)
5. Create new Pull Request

