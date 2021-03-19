# Dotnet Inner Loop Example App

## To build the app container
```bash
pack build dotnet-develop --descriptor project.toml --builder paketobuildpacks/builder:base --trust-builder
```

## To run the app container (and enable watching source-code changes)
```bash
docker run -it  -v "$(pwd)"/src:/workspace/src  dotnet-develop
```
Note that the container has the source directory on the host machine mounted
into `/workspace/src`, the location of the source code in the app container.
This will allow changes made on the host machine to be visible to the app
container.

## To see that the app is starting correctly:
The container should show output similar to:

```

Welcome to .NET Core!
---------------------
Learn more about .NET Core: https://aka.ms/dotnet-docs
Use 'dotnet --help' to see available commands or visit: https://aka.ms/dotnet-cli-docs

Telemetry
---------
The .NET Core tools collect usage data in order to help us improve your experience. It is collected by Microsoft and shared with the community. You can opt-out of telemetry by setting the DOTNET_CLI_TELEMETRY_OPTOUT environment variable to '1' or 'true' using your favorite shell.

Read more about .NET Core CLI Tools telemetry: https://aka.ms/dotnet-cli-telemetry

Configuring...
--------------
A command is running to populate your local package cache to improve restore speed and enable offline access. This command takes up to one minute to complete and only runs once.
Decompressing 100% 11690 ms
Expanding 100% 26190 ms

ASP.NET Core
------------
Successfully installed the ASP.NET Core HTTPS Development Certificate.
To trust the certificate run 'dotnet dev-certs https --trust' (Windows and macOS only). For establishing trust on other platforms refer to the platform specific documentation.
For more information on configuring HTTPS see https://go.microsoft.com/fwlink/?linkid=848054.
watch : Started
Hello World!
```

## To change the source code and watch a rebuild take place
While the container is still running, edit `Program.cs` on the host machine to
print a different message. Save the file.

In the container output, you should see feedback that `dotnet` has detected the
change. It will rebuild the app and print the new messgae.

