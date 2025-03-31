package io

import (
	"fmt"

	"github.com/slok/sloth/internal/info"
)

var yamlTopdisclaimer = fmt.Sprintf(`
---
# Code generated by Sloth (%s): https://github.com/slok/sloth.
# DO NOT EDIT.

`, info.Version)

func writeYAMLTopDisclaimer(bs []byte) []byte {
	return append([]byte(yamlTopdisclaimer), bs...)
}
