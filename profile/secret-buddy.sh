#!/usr/bin/env bash

PROFILE_D_DIR=/app/.profile.d
BUILDPACK=$PROFILE_D_DIR/secret-buddy-buildpack
($BUILDPACK) > $PROFILE_D_DIR/secret-buddy-env.sh
source $PROFILE_D_DIR/secret-buddy-env.sh