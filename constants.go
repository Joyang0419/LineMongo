package LineMongo

import (
	"LineMongo/configs"
	"os"
)

const (
	DateStringFormat = "2006-01-02"
)

var (
	path, _   = os.Getwd()
	Config, _ = configs.LoadConfig(path, "app", "env")
)
