#!/bin/bash

ytt -f config \
	-f config/_ytt_lib/cc/release/image.yml \
	-f config/_ytt_lib/uaa/release/image.yml \
	-f user-provided/values-with-db.yml
