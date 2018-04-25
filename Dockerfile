FROM plugins/base:amd64

LABEL maintainer="Drone.IO Community <drone-dev@googlegroups.com>" \
  org.label-schema.name="Drone Anynines" \
  org.label-schema.vendor="Drone.IO Community" \
  org.label-schema.schema-version="1.0"

RUN apk add --no-cache git ruby && \
  gem install --no-ri --no-rdoc dpl

ADD release/linux/amd64/drone-anynines /bin/
ENTRYPOINT ["/bin/drone-anynines"]
