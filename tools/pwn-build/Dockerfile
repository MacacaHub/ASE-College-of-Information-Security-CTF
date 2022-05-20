FROM ubuntu:20.04

RUN apt-get -qq update
RUN apt-get -qqy install gcc make

RUN useradd -M builder
USER builder

WORKDIR /src
