package main

import (
	"exc/internal/config"
)

func init() {
	//os.Setenv("CONFIG_PATH", "./internal/config/local.yaml ./your-app")


	config.MustLoad()
}
func main() {
}
