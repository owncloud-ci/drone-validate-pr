FROM docker.io/alpine:3.18@sha256:82d1e9d7ed48a7523bdebc18cf6290bdb97b82302a8a9c27d4fe885949ea94d1

LABEL maintainer="ownCloud DevOps <devops@owncloud.com>"
LABEL org.opencontainers.image.authors="ownCloud DevOps <devops@owncloud.com>"
LABEL org.opencontainers.image.title="drone-fork-approval"
LABEL org.opencontainers.image.url="https://github.com/owncloud-ci/drone-fork-approval"
LABEL org.opencontainers.image.source="https://github.com/owncloud-ci/drone-fork-approval"
LABEL org.opencontainers.image.documentation="https://github.com/owncloud-ci/drone-fork-approval"

ADD dist/drone-fork-approval /bin/

ENTRYPOINT ["/bin/drone-fork-approval"]
