FROM debian:stretch-slim

WORKDIR /

COPY sample-scheduler /usr/local/bin

ENTRYPOINT ["sample-scheduler"]