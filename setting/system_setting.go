package setting

import (
	"os"
)

var ServerAddress = getServerAddress()
var WorkerUrl = ""
var WorkerValidKey = ""
var WorkerAllowHttpImageRequestEnabled = false

func getServerAddress() string {
	if addr := os.Getenv("SERVER_ADDRESS"); addr != "" {
		return addr
	}
	return "http://localhost:3000"
}

func EnableWorker() bool {
	return WorkerUrl != ""
}
