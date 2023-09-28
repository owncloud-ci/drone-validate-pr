FROM docker.io/alpine:3.18@sha256:eece025e432126ce23f223450a0326fbebde39cdf496a85d8c016293fc851978

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
