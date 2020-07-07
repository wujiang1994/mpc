package main

import (
	"flag"
	"mpc"
	"mpc/testapp/app"
)

var (
	runMode string
	cfgPath string
)

func main() {
	flag.StringVar(&runMode, "runMode", "dev", "")
	flag.StringVar(&cfgPath, "cfgPath", "/testapp", "")
	flag.Parse()

	mpc.NewAppServer(runMode, cfgPath).NewServer(app.New()).Run()
}
