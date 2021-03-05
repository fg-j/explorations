package hugobuildpack_test

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"

	hugobuildpack "github.com/fg-j/explorations/hugo-buildpack"
	"github.com/fg-j/explorations/hugo-buildpack/fakes"
	"github.com/paketo-buildpacks/packit"
	"github.com/paketo-buildpacks/packit/postal"
	"github.com/sclevine/spec"

	. "github.com/onsi/gomega"
)

func testBuild(t *testing.T, context spec.G, it spec.S) {
	var (
		Expect = NewWithT(t).Expect

		layersDir         string
		workingDir        string
		cnbDir            string
		executable        *fakes.Executable
		entryResolver     *fakes.EntryResolver
		dependencyManager *fakes.DependencyManager

		build packit.BuildFunc
	)

	it.Before(func() {
		var err error
		layersDir, err = ioutil.TempDir("", "layers")
		Expect(err).NotTo(HaveOccurred())

		cnbDir, err = ioutil.TempDir("", "cnb")
		Expect(err).NotTo(HaveOccurred())

		workingDir, err = ioutil.TempDir("", "working-dir")
		Expect(err).NotTo(HaveOccurred())

		entryResolver = &fakes.EntryResolver{}
		entryResolver.ResolveCall.Returns.BuildpackPlanEntry = packit.BuildpackPlanEntry{
			Name: "hugo",
		}

		dependencyManager = &fakes.DependencyManager{}
		dependencyManager.ResolveCall.Returns.Dependency = postal.Dependency{
			ID:      "hugo",
			Name:    "hugo-dependency-name",
			SHA256:  "hugo-dependency-sha",
			Stacks:  []string{"some-stack"},
			URI:     "hugo-dependency-uri",
			Version: "hugo-dependency-version",
		}

		executable = &fakes.Executable{}

		build = hugobuildpack.Build(entryResolver, dependencyManager, executable)
	})

	it.After(func() {
		Expect(os.RemoveAll(layersDir)).To(Succeed())
		Expect(os.RemoveAll(cnbDir)).To(Succeed())
		Expect(os.RemoveAll(workingDir)).To(Succeed())
	})

	it("returns a result that installs hugo and builds correctly", func() {
		result, err := build(packit.BuildContext{
			WorkingDir: workingDir,
			CNBPath:    cnbDir,
			Stack:      "some-stack",
			BuildpackInfo: packit.BuildpackInfo{
				Name:    "Some Buildpack",
				Version: "some-version",
			},
			Plan: packit.BuildpackPlan{
				Entries: []packit.BuildpackPlanEntry{
					{
						Name: "hugo",
					},
				},
			},
			Layers: packit.Layers{Path: layersDir},
		})
		Expect(err).NotTo(HaveOccurred())

		Expect(result).To(Equal(packit.BuildResult{
			Plan: packit.BuildpackPlan{
				// Entries: []packit.BuildpackPlanEntry{
				// 	{
				// 		Name: "hugo",
				// 		Metadata: map[string]interface{}{
				// 			"name":   "hugo-dependency-name",
				// 			"sha256": "hugo-dependency-sha",
				// 			"stacks": []string{"some-stack"},
				// 			"uri":    "hugo-dependency-uri",
				// 		},
				// 	},
				// },
			},
			Layers: []packit.Layer{
				{
					Name:      "hugo",
					Path:      filepath.Join(layersDir, "hugo"),
					SharedEnv: packit.Environment{},
					BuildEnv:  packit.Environment{},
					LaunchEnv: packit.Environment{},
					Build:     false,
					Launch:    false,
					Cache:     false,
					// Metadata:  map[string]interface{}{
					// 	"dependency-sha": "hugo-dependency-sha",
					// 	"built_at":       timestamp.Format(time.RFC3339Nano),
					// },
				},
			},
		}))
	})
}
