#!/usr/bin/env bash

if [[ $1 == "set" ]];
then
echo "设置Go Proxy Shadowsock"
export http_proxy=http://127.0.0.1:1080
git config --global http.proxy http://127.0.0.1:1080
elif [[ $1 == "unset" ]];
then
echo "取消设置Go Proxy Shadowsock"
export http_proxy=''
git config --global --unset http.proxy
else
echo "unknown action"
fi
