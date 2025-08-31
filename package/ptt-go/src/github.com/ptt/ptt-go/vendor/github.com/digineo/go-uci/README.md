# go-uci

> **WORK IN PROGRESS**

[![GoDoc](https://pkg.go.dev/badge/github.com/digineo/go-uci.svg)](https://pkg.go.dev/github.com/digineo/go-uci)
[![Test results](https://github.com/digineo/go-uci/workflows/Test/badge.svg)](https://github.com/digineo/go-uci/actions?query=workflow%3ATest)
[![Lint results](https://github.com/digineo/go-uci/workflows/Lint/badge.svg)](https://github.com/digineo/go-uci/actions?query=workflow%3ALint)
[![Codecov](http://codecov.io/github/digineo/go-uci/coverage.svg?branch=master)](http://codecov.io/github/digineo/go-uci?branch=master)


UCI is OpenWrt's [Unified Configuration Interface][uci-wiki]. It is
used to configure OpenWrt router hardware using a simple DSL (and
acompanying CLI tools). Configuration files are written into a
central directory (`/etc/config/*`) which basically represents a
key/value store.

This project makes it easy to interact with such a config tree by
providing a native interface to that KV store. It has no external
runtime dependencies.

For now, we only implements a superset of the actual UCI DSL, but
improvements (patches or PRs) are very welcome. Refer to Rob Pike's
[Lexical Scanning in Go][pike-lex] for implementation details on the
parser/lexer.

[uci-wiki]: https://openwrt.org/docs/guide-user/base-system/uci
[pike-lex]: https://talks.golang.org/2011/lex.slide

## Why?

We're currently experimenting with Go binaries on OpenWrt router
hardware and need a way to interact with the system configuration.
We could have created bindings for [`libuci`][uci-git], but the
turnaround cycle in developing with CGO is a bit tedious. Also, since
Go does not compile for our target platforms, we need to resort to
GCCGO, which has other quirks.

The easiest solution therefore is a plain Go library, which can be
used in Go (with or without CGO) and GCCGO without worrying about
interoperability. A library also allows UCI to be used outside of
OpenWrt systems (e.g. for provisioning).

[uci-git]: https://git.openwrt.org/?p=project/uci.git;a=summary


## Usage

```go
import "github.com/digineo/go-uci"

func main() {
    // use the default tree (/etc/config)
    if values, ok := uci.Get("system", "@system[0]", "hostname"); ok {
        fmt.Println("hostanme", values)
        //=> hostname [OpenWrt]
    }

    // use a custom tree
    u := uci.NewTree("/path/to/config")
    if values, ok := u.Get("network", "lan", "ifname"); ok {
        fmt.Println("network.lan.ifname", values)
        //=> network.lan.ifname [eth0.2]
    }
    if sectionExists := u.Set("network", "lan", "ipaddr", "192.168.7.1"); !sectionExists {
        _ = u.AddSection("network", "lan", "interface")
        _ = u.Set("network", "lan", "ipaddr", "192.168.7.1")
    }
    u.Commit() // or uci.Revert()
}
```

See [API documentation][godoc] for more details.


## Contributing

Pull requests are welcome, especially if they increase test coverage.

Before submitting changes, please make sure the tests still pass:

```console
$ go test github.com/digineo/go-uci/...
```


## License

MIT License. Copyright (c) 2019 Dominik Menke, Digineo GmbH

<https://www.digineo.de>

See LICENSE file for details.
