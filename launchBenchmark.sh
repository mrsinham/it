#!/bin/bash
for goversion in 1.4.0 1.4.2 1.5.0 1.5.1; do sudo docker build -f $goversion -t interface_test:$go_version . && sudo docker run interface_test:$go_version ; done
