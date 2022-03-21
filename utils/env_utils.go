package utils

import "os"

var ServerHost string

func LazyEnvVariableInit() {

	ServerHost = os.Getenv("SERVER_HOST")
}
