# drone-fork-approval

[![Build Status](https://drone.owncloud.com/api/badges/owncloud-ci/drone-fork-approval/status.svg)](https://drone.owncloud.com/owncloud-ci/drone-fork-approval)
[![Docker Hub](https://img.shields.io/docker/v/owncloudci/drone-fork-approval?logo=docker&label=dockerhub&sort=semver&logoColor=white)](https://hub.docker.com/r/owncloudci/drone-fork-approval)
[![GitHub contributors](https://img.shields.io/github/contributors/owncloud-ci/drone-fork-approval)](https://github.com/owncloud-ci/drone-fork-approval/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/owncloud-ci/drone-fork-approval)
[![License: Apache-2.0](https://img.shields.io/github/license/owncloud-ci/drone-fork-approval)](https://github.com/owncloud-ci/drone-fork-approval/blob/main/LICENSE)

Drone CI validation extension to ensures that any PR originating from a fork must be manually approved before it is executed. This extension is a reimplementation of [drone-plugins/drone-docker](https://github.com/wadells/drone-fork-approval-extension).

## Usage

Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Download and run the plugin:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_FORK_APPROVAL_BIND=:3000 \
  --env=DRONE_FORK_APPROVAL_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --env=DRONE_FORK_APPROVAL_LOG_LEVEL=Info \
  --restart=always \
  --name=fork-approval owncloudci/drone-fork-approval
```

Update the Drone server configuration to include the plugin address and the shared secret.

```text
DRONE_VALIDATE_PLUGIN_ENDPOINT=http://127.0.0.1:3000
DRONE_VALIDATE_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6
```

## Build

Build the binary with the following command:

```console
make build
```

Build the Docker image with the following command:

```console
docker build --file Dockerfile.multiarch --tag owncloudci/drone-fork-approval .
```

## Releases

Create and push the new tag to trigger the CI release process:

```console
git tag v2.10.3
git push origin v2.10.3
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](https://github.com/owncloud-ci/drone-fork-approval/blob/main/LICENSE) file for details.

## Copyright

```text
Copyright (c) 2023 ownCloud GmbH
```
