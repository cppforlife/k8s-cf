#!/bin/bash

ytt -f config \
	-f config/_ytt_lib/cc/config/build.yml \
	-f config/_ytt_lib/uaa/config/build.yml \
	-f hack/build-from-source.yml \
	-f user-provided/values-with-db.yml
