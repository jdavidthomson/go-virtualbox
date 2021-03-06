# go-virtualbox

This is a wrapper package for Golang to interact with VirtualBox. The API is
experimental at the moment and you should expect frequent changes.

API doc at http://godoc.org/github.com/riobard/go-virtualbox

**Table of Contents**

<!-- TOC depthFrom:2 depthTo:4 -->

- [Status](#status)
- [Building](#building)
- [Testing](#testing)
    - [Preparation](#preparation)
    - [Run tests](#run-tests)

<!-- /TOC -->

## Status

[![CircleCI](https://circleci.com/gh/asnowfix/go-virtualbox.svg?style=svg)](https://circleci.com/gh/asnowfix/go-virtualbox)

| | Location |
|-|----------|
| CircleCI | https://circleci.com/gh/asnowfix/go-virtualbox |

## Building

First install dependencies

```bash
$ go get -v github.com/golang/dep/cmd/dep
$ dep ensure -v
```

Then build:

```bash
$ go build -v
```
## Testing 

### Preparation

Tests run using mocked interfaces, unless the `TEST_VM` environment variable is set, in order to test against real VirtualBox. You either need to  have a pre-provisioned VirtualBox VM and  to set its name using the `TEST_VM` environment variable, or use [Vagrant](https://www.vagrantup.com/intro/getting-started/).

```bash
$ vagrant box add bento/ubuntu-16.04
# select virtualbox as provider

$ vagrant up
```

Then run the tests 

```bash
$ export TEST_VM=go-virtualbox
$ go test
```

...or (on Windows):

```shell
> set TEST_VM=go-virtualbox
> go test
```

Once you are done with testing, run `vagrant halt` to same resources.

### Run tests

As usual, run `go test`, or `go test -v`.  To run one test in particular,
run `go test --run TestGuestProperty`.

