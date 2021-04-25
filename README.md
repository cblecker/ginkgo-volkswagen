# ginkgo-volkswagen

[![GoDoc](https://pkg.go.dev/badge/github.com/cblecker/ginkgo-volkswagen)](https://pkg.go.dev/github.com/cblecker/ginkgo-volkswagen)
[![Go](https://github.com/cblecker/ginkgo-volkswagen/actions/workflows/go.yml/badge.svg)](https://github.com/cblecker/ginkgo-volkswagen/actions/workflows/go.yml)

ginko-volkswagen detects when your tests are being run in a CI server, and reports them as passing

## Installation
```bash
go get github.com/cblecker/ginkgo-volkswagen
```

## Usage
In order to use this package, you need use the `Fail` function as your failure handler:
```go
import ginkgovolkswagen "github.com/cblecker/ginkgo-volkswagen"

func TestSuite(t *testing.T) {
    RegisterFailHandler(ginkgovolkswagen.Fail)
    ginkgo.RunSpecs(t, "Validation Tests")
}
```

Then enjoy your green CI results!

## Credits
Heavily inspired by https://github.com/auchenberg/volkswagen
