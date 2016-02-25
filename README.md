# drone-anynines

[![Build Status](http://beta.drone.io/api/badges/drone-plugins/drone-anynines/status.svg)](http://beta.drone.io/drone-plugins/drone-anynines)
[![Coverage Status](https://aircover.co/badges/drone-plugins/drone-anynines/coverage.svg)](https://aircover.co/drone-plugins/drone-anynines)
[![](https://badge.imagelayers.io/plugins/drone-anynines:latest.svg)](https://imagelayers.io/?images=plugins/drone-anynines:latest 'Get your own badge on imagelayers.io')

Drone plugin to deploy or update a project on Anynines. For the usage information and a listing of the available options please take a look at [the docs](DOCS.md).

## Binary

Build the binary using `make`:

```
make deps build
```

### Example

```sh
./drone-anynines <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
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

Build the container using `make`:

```
make deps docker
```

### Example

```sh
docker run -i plugins/drone-anynines <<EOF
{
    "repo": {
        "clone_url": "git://github.com/drone/drone",
        "owner": "drone",
        "name": "drone",
        "full_name": "drone/drone"
    },
    "system": {
        "link_url": "https://beta.drone.io"
    },
    "build": {
        "number": 22,
        "status": "success",
        "started_at": 1421029603,
        "finished_at": 1421029813,
        "message": "Update the Readme",
        "author": "johnsmith",
        "author_email": "john.smith@gmail.com"
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
