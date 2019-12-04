FROM ubuntu:latest

#Base Info
MAINTAINER jeferwang@qq.com

#System Env Config
RUN sed -i 's/archive.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN sed -i 's/security.ubuntu.com/mirrors.ustc.edu.cn/g' /etc/apt/sources.list
RUN apt-get update
RUN apt-get install -y apt-utils curl

#Installation
ENV ETCD_VER=v3.3.17
ARG GOOGLE_URL=https://storage.googleapis.com/etcd
ARG GITHUB_URL=https://github.com/etcd-io/etcd/releases/download
ARG DOWNLOAD_URL=${GITHUB_URL}
RUN rm -rf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
RUN rm -rf /tmp/etcd-download-test
RUN mkdir -p /tmp/etcd-download-test
RUN export http_proxy=http://192.168.0.112:1080/ && export https_proxy=http://192.168.0.112:1080/ && curl -L ${DOWNLOAD_URL}/${ETCD_VER}/etcd-${ETCD_VER}-linux-amd64.tar.gz -o /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
RUN tar -xvf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz -C /tmp/etcd-download-test --strip-components=1
RUN rm -rf /tmp/etcd-${ETCD_VER}-linux-amd64.tar.gz
RUN cp -r /tmp/etcd-download-test/etcd /usr/local/bin/etcd
RUN cp -r /tmp/etcd-download-test/etcdctl /usr/local/bin/etcdctl
RUN rm -rf /tmp/etcd-download-test

#Installation Test
RUN etcd --version
RUN etcdctl --version

#Use
ENV ETCDCTL_API=3
