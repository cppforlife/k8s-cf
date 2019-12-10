#!/bin/bash

kapp deploy -a cf -f <(./hack/template-prebuilt.sh | kbld -f-) -c
