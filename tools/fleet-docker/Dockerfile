FROM debian:stable-slim AS builder

RUN apt-get update
RUN apt-get install -y build-essential autoconf libxml2-dev libssl-dev zlib1g-dev curl

# Install bomutils
RUN curl -L https://github.com/hogliux/bomutils/archive/0.2.tar.gz > bomutils.tar.gz && \
    echo "fb1f4ae37045eaa034ddd921ef6e16fb961e95f0364e5d76c9867bc8b92eb8a4  bomutils.tar.gz" | sha256sum --check && \
    tar -xzf bomutils.tar.gz
RUN cd bomutils-0.2 && make && make install

# Install xar
RUN curl -L https://github.com/mackyle/xar/archive/refs/tags/xar-1.6.1.tar.gz > xar.tar.gz && \
    echo "5e7d50dab73f5cb1713b49fa67c455c2a0dd2b0a7770cbc81b675e21f6210e25  xar.tar.gz" | sha256sum --check && \
    tar -xzf xar.tar.gz 
# Note this needs patching due to newer version of OpenSSL
# See https://github.com/mackyle/xar/pull/23
COPY patch.txt .
RUN cd xar-xar-1.6.1/xar && patch < ../../patch.txt && autoconf && ./configure && make && make install

FROM debian:stable-slim

COPY --from=builder /usr/bin/mkbom /usr/local/bin/xar /usr/bin/
COPY --from=builder /usr/local/lib /usr/local/lib/
COPY make-aliases.sh /home/wine/make-aliases.sh

WORKDIR /home/wine
ENV HOME=/home/wine WINEPREFIX=/home/wine/.wine WINEARCH=win32 PATH="/home/wine/bin:$PATH" WINEDEBUG=-all

RUN true \
    && dpkg --add-architecture i386 \
    && apt update \
    && apt install -y --no-install-recommends \
    ca-certificates \
    wine \
    wine32 \
    wget \
    unzip \
    && mkdir /wix \
    && rm -rf /var/lib/apt/lists/*  \
    && wine wineboot \
    && wget https://dl.winehq.org/wine/wine-mono/6.4.0/wine-mono-6.4.0-x86.msi -nv -O mono.msi \
    && wget https://github.com/wixtoolset/wix3/releases/download/wix3112rtm/wix311-binaries.zip -nv -O wix.zip \
    && mkdir wix \
    && unzip wix.zip -d wix \
    && wine msiexec /i mono.msi \
    && /home/wine/make-aliases.sh \
    && rm -f /home/wine/make-aliases.sh mono.msi wix.zip \
    && mkdir $WINEPREFIX/drive_c/temp

