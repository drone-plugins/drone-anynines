# drone-anynines

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-anynines/status.svg)](http://beta.drone.io/drone-plugins/drone-anynines)
[![Join the discussion at https://discourse.drone.io](https://img.shields.io/badge/discourse-forum-orange.svg)](https://discourse.drone.io)
[![Drone questions at https://stackoverflow.com](https://img.shields.io/badge/drone-stackoverflow-orange.svg)](https://stackoverflow.com/questions/tagged/drone.io)
[![Go Doc](https://godoc.org/github.com/drone-plugins/drone-anynines?status.svg)](http://godoc.org/github.com/drone-plugins/drone-anynines)
[![Go Report](https://goreportcard.com/badge/github.com/drone-plugins/drone-anynines)](https://goreportcard.com/report/github.com/drone-plugins/drone-anynines)
[![](https://images.microbadger.com/badges/image/plugins/anynines.svg)](https://microbadger.com/images/plugins/anynines "Get your own image badge on microbadger.com")

Drone plugin to deploy or update a project on [Anynines](https://www.anynines.com/). For the usage information and a listing of the available options please take a look at [the docs](http://plugins.drone.io/drone-plugins/drone-anynines/).

## Build

Build the binary with the following commands:

```
go build
```

## Docker

Build the Docker image with the following commands:

```
GOOS=linux GOARCH=amd64 CGO_ENABLED=0 go build -a -tags netgo -o release/linux/amd64/drone-anynines
docker build --rm -t plugins/anynines .
```

### Usage

```
docker run --rm \
  -e PLUGIN_USERNAME=octocat@github.com \
  -e PLUGIN_PASSWORD=password \
  -e PLUGIN_ORGANIZATION=octocat_github_com \
  -e PLUGIN_SPACE=production \
  -v $(pwd):$(pwd) \
  -w $(pwd) \
  plugins/anynines
```
