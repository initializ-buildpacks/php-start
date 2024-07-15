package main

import (
	"os"

	phpstart "github.com/initializ-buildpacks/php-start"
	"github.com/paketo-buildpacks/libreload-packit/watchexec"
	"github.com/paketo-buildpacks/packit/v2"
	"github.com/paketo-buildpacks/packit/v2/scribe"
)

func main() {
	procMgr := phpstart.NewProcs()
	logEmitter := scribe.NewEmitter(os.Stdout).WithLevel(os.Getenv("BP_LOG_LEVEL"))

	reloader := watchexec.NewWatchexecReloader()

	packit.Run(
		phpstart.Detect(reloader),
		phpstart.Build(procMgr, logEmitter, reloader),
	)
}
