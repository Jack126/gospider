package bootstrap

import (
	"gospider/global/variable"
	"os"
)

func init() {
	// init base path
	path, _ := os.Getwd()
	variable.BasePath = path
}
