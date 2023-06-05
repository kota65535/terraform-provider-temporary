package provider

import (
	"path/filepath"
	"strings"
)

func ContainsFilePath(p, sub string) (bool, error) {
	pa, err := filepath.Abs(p)
	if err != nil {
		return false, err
	}
	suba, err := filepath.Abs(sub)
	if err != nil {
		return false, err
	}
	return strings.Contains(suba, pa), nil
}
