#!/usr/bin/env sh

yum install -y wget
yum install -y ruby
yum install -y rubygems

gem install mustache

sass_tar=dart-sass-1.21.0-linux-x64.tar.gz
wget https://github.com/sass/dart-sass/releases/download/1.21.0/$sass_tar
tar -xzf $sass_tar
export PATH="$(pwd)/dart-sass:$PATH"

sh build.sh dist
