FROM debian:stable-slim

RUN apt-get update && apt-get install -y --no-install-recommends libxml2  && rm -rf /var/lib/apt/lists/*

# copy macOS dependencies
COPY --from=fleetdm/bomutils:latest /usr/bin/mkbom /usr/local/bin/xar /usr/bin/
COPY --from=fleetdm/bomutils:latest /usr/local/lib /usr/local/lib/

# copy Windows dependencies
COPY --from=fleetdm/wix:latest /usr/bin/wine* /usr/bin/
COPY --from=fleetdm/wix:latest /home/wine /home/wine

# copy fleetctl
COPY build/binary-bundle/linux/fleetctl /usr/bin/fleetctl

RUN apt-get update && apt-get install -y ca-certificates cpio && rm -rf /var/lib/apt/lists/* 
