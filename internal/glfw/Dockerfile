FROM debian:testing

RUN apt-get update && apt-get install -y \
        ca-certificates \
        golang \
        gcc-mingw-w64 \
        && rm -rf /var/lib/apt/lists/*

WORKDIR /work
