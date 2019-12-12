#!/bin/bash

ytt -f config \
	-f hack/build-from-source.yml \
	-f user-provided/values-with-db.yml
