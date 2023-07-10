# drone-fork-approval

[![Build Status](https://img.shields.io/drone/build/owncloud-ci/drone-fork-approval?logo=drone&server=https%3A%2F%2Fdrone.owncloud.com)](https://drone.owncloud.com/owncloud-ci/drone-fork-approval)
[![Docker Hub](https://img.shields.io/docker/v/owncloudci/drone-fork-approval?logo=docker&label=dockerhub&sort=semver&logoColor=white)](https://hub.docker.com/r/owncloudci/drone-fork-approval)
[![GitHub contributors](https://img.shields.io/github/contributors/owncloud-ci/drone-fork-approval)](https://github.com/owncloud-ci/drone-fork-approval/graphs/contributors)
[![Source: GitHub](https://img.shields.io/badge/source-github-blue.svg?logo=github&logoColor=white)](https://github.com/owncloud-ci/drone-fork-approval)
[![License: Apache-2.0](https://img.shields.io/github/license/owncloud-ci/drone-fork-approval)](https://github.com/owncloud-ci/drone-fork-approval/blob/main/LICENSE)

Fork approval is a simple Drone CI validation extension that ensures that any PR originating from a fork must be manually approved before it is executed.

## Installation

Create a shared secret:

```console
$ openssl rand -hex 16
bea26a2221fd8090ea38720fc445eca6
```

Download and run the plugin:

```console
$ docker run -d \
  --publish=3000:3000 \
  --env=DRONE_DEBUG=true \
  --env=DRONE_SECRET=bea26a2221fd8090ea38720fc445eca6 \
  --restart=always \
  --name=starlark owncloudci/drone-fork-approval
```

Update your Drone server configuration to include the plugin address and the shared secret.

```text
DRONE_VALIDATE_PLUGIN_ENDPOINT=http://1.2.3.4:3000
DRONE_VALIDATE_PLUGIN_SECRET=bea26a2221fd8090ea38720fc445eca6
```

## License

This project is licensed under the Apache 2.0 License - see the [LICENSE](https://github.com/owncloud-ci/drone-fork-approval/blob/main/LICENSE) file for details.

## Copyright

```Text
Copyright (c) 2023 ownCloud GmbH
```
