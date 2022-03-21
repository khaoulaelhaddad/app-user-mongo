package utils

import (
	"os"
)

var ServerHost string
var ApplyURI string

func LazyEnvVariableInit() {

	ServerHost = os.Getenv("SERVER_HOST")
	ApplyURI = os.Getenv("APPLY_URI")
}
