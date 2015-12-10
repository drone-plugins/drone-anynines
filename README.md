# drone-anynines

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-anynines/status.svg)](http://beta.drone.io/drone-plugins/drone-anynines)
[![](https://badge.imagelayers.io/plugins/drone-anynines:latest.svg)](https://imagelayers.io/?images=plugins/drone-anynines:latest 'Get your own badge on imagelayers.io')

Drone plugin for deploying to Anynines

## Usage

```
./drone-anynines <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "username": "octocat@github.com",
        "password": "password",
        "organization": "octocat_github_com",
        "space": "production"
    }
}
EOF
```

## Docker

Build the Docker container using `make`:

```
make deps build
docker build --rm=true -t plugins/drone-anynines .
```

### Example

```sh
docker run -i plugins/drone-anynines <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "full_name": "drone/drone"
    },
    "build": {
        "event": "push",
        "branch": "master",
        "commit": "436b7a6e2abaddfd35740527353e78a227ddcb2c",
        "ref": "refs/heads/master"
    },
    "workspace": {
        "root": "/drone/src",
        "path": "/drone/src/github.com/drone/drone"
    },
    "vargs": {
        "username": "octocat@github.com",
        "password": "password",
        "organization": "octocat_github_com",
        "space": "production"
    }
}
EOF
```
