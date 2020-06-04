package bootstrap

import (
	"log"
	"os"
	_ "spider/global/destroy"
	SpiderError "spider/global/error"
	"test/global/variable"
)

func init() {
	// init base path
	if path, err := os.Getwd(); err != nil {
		variable.BasePath = path
	} else {
		log.Fatal(SpiderError.ErrorsBasePath)
	}

}
