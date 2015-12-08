# Docker image for the Drone Anynines plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-anynines
#     GO15VENDOREXPERIMENT=1 CGO_ENABLED=0 go build -a -tags netgo
#     docker build --rm=true -t plugins/drone-anynines .

FROM alpine:3.2

RUN apk update && \
  apk add \
    ca-certificates \
    ruby && \
  gem install --no-ri --no-rdoc \
    dpl && \
  rm -rf /var/cache/apk/*

ADD drone-anynines /bin/
ENTRYPOINT ["/bin/drone-anynines"]
