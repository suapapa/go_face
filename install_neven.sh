#!/bin/bash

git submodule init
git submodule update
pushd neven
scons
sudo scons install
popd
