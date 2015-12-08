# Docker image for the Drone Anynines plugin
#
#     cd $GOPATH/src/github.com/drone-plugins/drone-anynines
#     make deps build
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
