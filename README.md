# addr

![License](https://img.shields.io/badge/license-Apache2.0-green) ![Language](https://img.shields.io/badge/Language-Go-blue.svg)

Addr is a simple library define Address struct, and support parsing from string or environment variable.

```
type Address struct {
	// Host domain name or ipv4, ipv6 address
	Host string
	// Port service port
	Port int
}
```

## 📋 Requirements

- Go 1.17+

## 🚀 Install

```
go get github.com/libgox/addr
```
