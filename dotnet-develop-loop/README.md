# Creating an "Inner Development Loop" With Existing .NET Buildpacks

The [sample app](sample-app/) uses existing buildpacks to create a running
container whose source code is still available. By volume mounting the `src`
directory on your local machine to the `src` directory in the container, it's
possible to set up an "inner loop" that relies on the `dotnet watch run`
command to watch for changes and rebuild the app.

### Unresolved problems and Hacks

#### Using the build plan buildpack
I had to use the build plan buildpack to force the relevant .NET dependencies
to be present in the app container. In the typical working of the [Paketo .NET
Core Buildpack](https://github.com/paketo-buildpacks/dotnet-core), the Dotnet
Publish and Dotnet Execute buildpacks are responsible for adding requirements
to the Build Plan. (See architecture
[RFC](https://github.com/paketo-buildpacks/dotnet-core/blob/main/rfcs/0001-restructure.md)
for details.)

In addition, because the Dotnet Publish and Dotnet Execute buildpacks are
typically responsible for parsing the project file to discover needed versions
of the .NET runtime and other dependencies, excluding the buildpacks from the
order group means that the correct version requirement must be made in the
`plan.toml`. However, in order for the Dotnet Runtime Buildpack to prioritize
the version requirement in the build plan, a `version-source` must be added:

```toml
[[ requires ]]
  name = "dotnet-runtime"

  [requires.metadata]
    launch=true
    build=true
    version = "2.1.*"
    version-source = "console.csproj"
```

This very hacky. It requires the user to know a lot about the of the Dotnet
Runtime Buildpack's [Plan Entry
resolver](https://github.com/paketo-buildpacks/dotnet-core-runtime/blob/main/plan_entry_resolver.go).

#### Using Procfile
Since the Dotnet Execute buildpack was excluded from the build order, the
buildpack had no default process. I used a Procfile to run `cd src && dotnet
watch run`. This is a bit brittle because it requires a user to know that the
working directory of the launch process will contain their app.

#### What happens if the dotnet dependencies need to change their version?
If the project file is changed to require a verison of the .NET runtime that
the buildpack has not installed, `dotnet watch run` will fail after attempting
to rebuild the app, because the needed runtime is not available in the
container.  The error is informative enough that a user could probably figure
out that they need to re-run their build.

### Takeaways
All in all, it wasn't that hard to generate a container that had some ability
to rebuild its app upon changes to the source code.

Adding a buildpack that is capable of running `dotnet watch run` *and* that has
similar detection logic as the Dotnet Execute and Dotnet Publish buildpacks
could do a lot to make this less hacky.  The buildpack could parse project
files, etc. to make proper version requirements and set the `dotnet watch run`
start command. This would eliminate the need for the Build Plan and Procfile
buildpacks.

### Remaining Questions
Do other language families have similar commands that watch source and rebuild
it? We might be able to leverage those and create start command buildpacks that
set up source code watching AND make the appropriate run-time dependency
requirements along the way.

