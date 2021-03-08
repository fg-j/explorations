package hugobuildpack

import (
	"path/filepath"

	"github.com/paketo-buildpacks/packit"
)

func Detect() packit.DetectFunc {
	return func(context packit.DetectContext) (packit.DetectResult, error) {
		files, err := filepath.Glob(filepath.Join(context.WorkingDir, "content", "*.md"))
		if err != nil {
			return packit.DetectResult{}, err
		}

		html, err := filepath.Glob(filepath.Join(context.WorkingDir, "content", "*.html"))
		if err != nil {
			panic(err)
			return packit.DetectResult{}, err
		}

		files = append(files, html...)
		if len(files) == 0 {
			return packit.DetectResult{}, packit.Fail.WithMessage("no *.md or *.html files found in content dir")
		}

		return packit.DetectResult{
			Plan: packit.BuildPlan{
				Requires: []packit.BuildPlanRequirement{
					{
						Name: "hugo",
						Metadata: map[string]interface{}{
							"build": true,
						},
					},
					{
						Name: "httpd",
						Metadata: map[string]interface{}{
							"launch": true,
						},
					},
				},
				Provides: []packit.BuildPlanProvision{
					{
						Name: "hugo",
					},
				},
			},
		}, nil
	}
}
