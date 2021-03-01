# Hugo Buildpack

## Detection
According to Hugo's [documentation about typical directory structure](https://gohugo.io/getting-started/directory-structure/), hugo projects tend to have directories:
```
archetypes/
assets/
config/ | config.toml
content/
data/
layouts/
static/
resources/
```

Where **content** is probably the biggest Hugo tipoff, because that's the
directory under which Hugo expects to find all of your site's pages. **config**
also seems pertinent. Things like **static** won't be as helpful because they
are common to many frameworks.

According to the Hugo [documentation about Content
Formats](https://gohugo.io/content-management/formats/), the content directory
will probably contain at least some `*.md` or `.html`. There are a few
generators/text processors that Hugo supports.  These are listed in the [list
of content
formats](https://gohugo.io/content-management/formats/#list-of-content-formats).These
include things like [Asciidoctor](https://asciidoctor.org/) (which converts
plain text to HTML). There are certain file extensions that signal each of
these processor types.

Looking for files that need to be processed with Hugo's supported processors
and installing those tools would be part of robust Hugo support in a buildpack.
This merits further investigation.

### Detecting Asciidoctor dependency

According to the [Asciidoctor
documentation](https://docs.asciidoctor.org/asciidoc/latest/document-structure/#encodings-and-asciidoc-files),
ASCIIdoctor processes files ending in `*.adoc`

Finding files in `content/**/*.adoc` would suggest Hugo needs ASCIIdoctor
installed.

### Detecting Emacs Org-mode (does this require a dependency?)

### Detecting RST Seems like this is a thing that relies on Python to be
converted to HTML. Expect the docs to be `*.py` files that include something
like: ``` __docformat__ = 'restructuredtext' ```

This seems like A Whole Can of Worms.

### Detecting Pandoc


## Build

When you want to build your hugo site, you [run `hugo` from within the
directory](https://gohugo.io/getting-started/usage/#deploy-your-website) of the
hugo site. This generates static HTML for the site and puts them all in the
`public/` directory. We will **not** want to preserve the public directory
between builds.


Seems like Hugo integrates with various ways of serving static sites. You can
even use the hugo CLI directory to deploy to AWS, Azure, Gcloud.  There's some
[deployment
config](https://gohugo.io/hosting-and-deployment/hugo-deploy/#configure-the-deployment)
that indicates how much cached content to save and so forth.

We would probably want to require some sort of buildpack that actually serves
the assets. These are the Web Servers buildpacks:
- [staticfile](https://github.com/paketo-community/staticfile) +
  [NGINX](https://github.com/paketo-buildpacks/nginx)
- [HTTPD](https://github.com/paketo-buildpacks/httpd)

The Hugo buildpack could require either of those. It can provide the hugo
distribution I guess (and also require it? Is that a weird pattern?).

There could be one buildpack that installs hugo and another that runs the
deploy I guess.  Hugo dependency buildpack just installs the binary (provides
Hugo) and the hugo-generate buildpack actually generates the files Requires
hugo binary, also requires other binaries or runtimes to generate various other
types of files.

## Layer caching When you run `hugo`, does it place things in `resources`?

It seems that when you run `hugo` to generate the site, some cached content is written to:
```
--cacheDir string            filesystem path to cache directory. Defaults: $TMPDIR/hugo_cache/
```
This could be specified as a layer location to speed rebuilds.


## Open Questions
- Seems like Hugo has lots of integrations to avoid your actually having to run
  your Hugo site from your own OCI image. Is there really a use-case for
  building a Hugo site with buildpacks?
- What's the proper way for a hugo deploy buildpack to require a subsequent web
  server?
- Can you configure a Hugo build with env vars? Or only with flags? If only
  with flags, how best to allow people to pass their desired flags through the
  to the build?
- What's the story with Hugo modules? (Seems like a new feature) and should we
  be running a `hugo mod vendor` before we do a build?
