#!/bin/bash

set -e -u

fly -t c2c \
  set-pipeline -p cf-networking-v1.10.x \
  -c $HOME/workspace/cf-networking-release/ci/pipelines/cf-networking.yml \
  -l $HOME/workspace/cf-networking-deployments/pipeline-credentials.yml
