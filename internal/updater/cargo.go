package updater

import (
	"regexp"
	"strings"
)

// Cargo creates an updater that modifies the version field in Cargo.toml files
func Cargo() Updater {
	return cargo{}
}

type cargo struct{}

func (c cargo) Files() []string {
	return []string{"Cargo.toml"}
}

func (c cargo) CreateNewFiles() bool {
	return false
}

func (c cargo) Update(info ReleaseInfo) func(content string) (string, error) {
	return func(content string) (string, error) {
		version := strings.TrimPrefix(info.Version, "v")

		versionRegex := regexp.MustCompile(`(version\s*=\s*)"[^"]*"`)

		if !versionRegex.MatchString(content) {
			return content, nil
		}

		return versionRegex.ReplaceAllString(content, `${1}"`+version+`"`), nil
	}
}
