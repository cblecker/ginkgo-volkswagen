// Package ginkgovolkswagen ensures that your tests always pass in CI :)
package ginkgovolkswagen

import (
	"os"

	"github.com/onsi/ginkgo/v2"
)

var ciEnvVariables = []string{
	"TRAVIS",
	"CI",
	"CONTINUOUS_INTEGRATION",
}

func isThisCI() bool {
	for _, variable := range ciEnvVariables {
		if _, set := os.LookupEnv(variable); set {
			return true
		}
	}

	return false
}

//Fail notifies Ginkgo that the current spec has failed.
func Fail(message string, callerSkip ...int) {
	if !isThisCI() {
		ginkgo.Fail(message, callerSkip...)
	}
}
