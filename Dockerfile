FROM docker.io/alpine:3.18@sha256:7144f7bab3d4c2648d7e59409f15ec52a18006a128c733fcff20d3a4a54ba44a

LABEL maintainer="ownCloud DevOps <devops@owncloud.com>"
LABEL org.opencontainers.image.authors="ownCloud DevOps <devops@owncloud.com>"
LABEL org.opencontainers.image.title="drone-fork-approval"
LABEL org.opencontainers.image.url="https://github.com/owncloud-ci/drone-fork-approval"
LABEL org.opencontainers.image.source="https://github.com/owncloud-ci/drone-fork-approval"
LABEL org.opencontainers.image.documentation="https://github.com/owncloud-ci/drone-fork-approval"

ADD dist/drone-fork-approval /bin/

RUN addgroup -g 1001 -S app && \
    adduser -S -D -H -u 1001 -s /sbin/nologin -G app -g app app

RUN apk --update add --no-cache ca-certificates && \
    rm -rf /var/cache/apk/* && \
    rm -rf /tmp/*

EXPOSE 3000

USER app

ENTRYPOINT ["/bin/drone-fork-approval"]
CMD []
