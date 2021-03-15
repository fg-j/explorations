# Layer Reuse in Dotnet Builds

There are two potential opportunities for layer reuse in the dotnet build process:
1. Reuse nuget packages from previous builds (Similar to having a `node_modules` layer
1. Restore built artifacts from previous builds (`*.dll` files and similar)

# NuGet Package Reuse
`dotnet publish` implicitly includes a call to `dotnet restore` (unless the
`dotnet publish --no-restore` flag is used). `dotnet restore` reaches out to
the internet and installs NuGet packages onto the host machine. These are, by
default, installed in the global packages folder located at `~/.nuget/packages`
on Linux.  This location can be overridden with the `$NUGET_PACKAGES`
environment variable.

To allow NuGet package restoration on rebuilds, the buildpack can set
`$NUGET_PACKAGES` to a layer location during both `dotnet restore` and `dotnet
publish` steps. This can also be accomplished by placing a `packages.config`
file in the app directory and specifying a `repositoryPath` in that directory.
See [nuget.config
reference](https://docs.microsoft.com/en-us/nuget/reference/nuget-config-file#config-section)
for more information on this.  _Complicating factor: if we wanted to separate
`dotnet restore` and `dotnet publish` into two separate buildpacks, following
the pattern of `go mod vendor` and `go build`, the location of the NuGet
packages would need to be communicated from one buildpack to the other. This
makes the `packages.config` approach more viable in this case._

If NuGet packages were placed in a layer, that layer would need to be available
to rebuilds (i.e. `cache=true`) but would **not** need to be available at
launch time (i.e. `launch=false`) because the [`dotnet publish`
Description](https://docs.microsoft.com/en-us/dotnet/core/tools/dotnet-publish#description)
states that the publish output contains "[t]he application's dependencies,
which are copied from the NuGet cache into the output folder." That is, `dotnet
publish` will handle moving NuGet packages to the needed location as long as
they are available at build time.

### Open question: Does restoring NuGet packages significantly improve rebuild times?

### Open question: Do we see packages in that global default location in a containerized app with buildpacks today?

### Open question: Should `dotnet restore` and `dotnet publish` be separated into two buildpacks?

### Open question: Are there advantages to `dotnet restore --use-lock-file`?

In [this blog
post](https://medium.com/01001101/containerize-your-net-core-app-the-right-way-35c267224a8d)
about containerizing dotnet apps, the writer recommends separating the `dotnet
restore` and `dotnet publish` so that the results can be built into separate
layers. The author claims that this improves rebuild speeds.

## Reusing built artifacts
As far as I can tell, `dotnet publish` does not natively cache anything to speed rebuilds.

### Open question: Does benchmark testing building/rebuilding the same app seem to speed up on rebuilds?

We could still place the `dotnet publish` output in a layer and reuse it under
some conditions. The obvious first choice is that if no source code changes,
the app need not be rebuilt.  But this is probably not true assuming that the
installed version of `dotnet` has changed. That is, if source code has not
changed but the installed versions of the dotnet runtime, SDK, or ASP.Net have
changed, we will probably want to rebuild the app.

### Open question: Can we safely reuse built artifacts even if the underlying runtime, SDK or Aspnet versions change?
