#!/bin/bash

#D:/Programs/Git/bin/bash.exe

protoc --proto_path=proto --micro_out=proto/ --go_out=proto/ proto/*.proto
