# addr

![License](https://img.shields.io/badge/license-Apache2.0-green)
![Language](https://img.shields.io/badge/Language-Go-blue.svg)
[![version](https://img.shields.io/github/v/tag/libgox/addr?label=release&color=blue)](https://github.com/libgox/addr/releases)
[![Godoc](http://img.shields.io/badge/docs-go.dev-blue.svg?style=flat-square)](https://pkg.go.dev/github.com/libgox/addr)
[![Go report](https://goreportcard.com/badge/github.com/libgox/addr)](https://goreportcard.com/report/github.com/libgox/addr)
[![codecov](https://codecov.io/gh/libgox/addr/branch/main/graph/badge.svg)](https://codecov.io/gh/libgox/addr)

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
