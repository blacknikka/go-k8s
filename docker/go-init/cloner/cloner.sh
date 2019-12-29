#!/bin/bash
set -Ceu

if [ -z $CONTENTS_SOURCE_URL ]; then
    exit 1
fi

git clone $CONTENTS_SOURCE_URL
