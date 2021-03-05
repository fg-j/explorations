package main

import (
	hugobuildpack "github.com/fg-j/explorations/hugo-buildpack"
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/cargo"
	"github.com/paketo-buildpacks/packit/draft"
	"github.com/paketo-buildpacks/packit/pexec"
	"github.com/paketo-buildpacks/packit/postal"
)

func main() {
	entryResolver := draft.NewPlanner()
	dependencyManager := postal.NewService(cargo.NewTransport())
	packit.Run(hugobuildpack.Detect(),
		hugobuildpack.Build(entryResolver, dependencyManager, pexec.NewExecutable("hugo")),
	)
}
