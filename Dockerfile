#imcomplete. still working on this.
FROM ubuntu:16.04
MAINTAINER Daiki Nakashita <daikinakashita.work@gmail.com>

# system dependencies
RUN apt-get update
RUN apt-get -y install bzr build-essential
RUN apt-get -y install curl git byobu
RUN apt-get -y install htop man unzip
RUN apt-get -y install vim wget mercurial pkg-config
RUN apt-get -y install software-properties-common

# go 1.7
RUN curl https://storage.googleapis.com/golang/go1.7.1.linux-amd64.tar.gz -O
RUN tar -C /usr/local -xzf go1.7.1.linux-amd64.tar.gz
ENV GOPATH /go
ENV HOME /go
