#!/bin/zsh
for goversion in 1.4.0 1.4.2 1.5.0 1.5.1 1.5.3 1.6.0 1.7.0; do echo ${goversion} && sudo docker build -f $goversion -t interface_test:${goversion} . && sudo docker run interface_test:${goversion} ; done
