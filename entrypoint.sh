#!/bin/sh

set -e

if [ "$1" = "app" ]; then
  RT_ARGS="-merlin-host $MERLIN_HOST -user $MERLIN_USER -pass $MERLIN_PASS"
fi

exec $@ $RT_ARGS
