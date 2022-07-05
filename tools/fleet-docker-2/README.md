## fleet/docker

This image needs to be built from the root of the repo so it has access to the build files. To build the image, run:

```
$ make xp-fleetctl
$ docker build --platform=linux/amd64 -f tools/fleet-docker-2/Dockerfile .
```
