package latest

import "io"

type Getter interface {
	GetLatestVersion(reader io.Reader, versionCondition string) (string, error)
}
