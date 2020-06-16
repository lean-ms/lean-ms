package helpers

import (
	"os"
	"path"
)

func GetBasePath() string {
	basepath, _ := os.Executable()
	return path.Join(basepath, "..", "..", "src", "github.com", "lean-ms", "lean-ms")
}
