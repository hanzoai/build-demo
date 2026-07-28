# build-demo

The smallest proof that `hanzo.app` builds software, not just container images.

`hanzo.yml` declares two artifacts and no Dockerfile:

- **`hanzo-demo`** — a static Go binary, cross-built for `linux/amd64` and `darwin/arm64`
- **`@hanzo/build-demo`** — a real npm package: `tsc` compiles `src/`, `npm pack` tars it

Both are built by `POST /v1/runner` on the platform's own cluster and published to
`https://s3.hanzo.ai/plugins/hanzoai/build-demo/<tag>/`, with a `binaries.json`
index carrying the URL and SHA-256 of each file.
