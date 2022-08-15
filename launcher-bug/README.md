# Investigation: Launcher unset default process bug

First, build the sample app with the sample buildpack that is checked in here.
This should work properly.

To build:
```
pack build test-sample --path ./sample-app --buildpack ./sample-buildpack --builder cnbs/sample-builder:bionic --clear-cache
```

To run:
```
docker run --rm -it test-sample
```

The expected output is:
```
Default entrypoint process
```

Then, modify `sample-buildpack/bin/build`. Swap out the working start process
block for the broken one, as indicated in the build script. Build and run using
the same commands. Expected output from run:
```
ERROR: failed to launch: determine start command: when there is no default process a command is required
```
