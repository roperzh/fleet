FROM debian:stable-slim

RUN apt-get update \
  && dpkg --add-architecture i386 \
  && apt update \
  && apt install -y --no-install-recommends ca-certificates cpio libxml2 wine wine32 \
  && rm -rf /var/lib/apt/lists/* 

# copy macOS dependencies
COPY --from=fleetdm/bomutils:latest /usr/bin/mkbom /usr/local/bin/xar /usr/bin/
COPY --from=fleetdm/bomutils:latest /usr/local/lib /usr/local/lib/

# copy Windows dependencies
COPY --from=fleetdm/wix:latest /home/wine /home/wine

# copy fleetctl
COPY build/binary-bundle/linux/fleetctl /usr/bin/fleetctl

ENV WINEPREFIX=/home/wine/.wine WINEARCH=win32 PATH="/home/wine/bin:$PATH" WINEDEBUG=-all FLEETCTL_NATIVE_TOOLING=1

ENTRYPOINT ["fleetctl"]
