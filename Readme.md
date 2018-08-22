# spaceoddity

## Overview

Spaceboy is a simple go  app for demoing K8s features:
* cpu intensive workload
* memory leak appllication

## Options

You can overwrite default parameters:
* "-cpu" bool (default false): cpu intensive workload
* "-mem" bool (default false): memory leak appllication
* "-second|-s" int: time in seconds

## Examples

`$ docker run -d de13/spaceoddity:v0.1 -cpu -s 300` # run intensite CPU app during 5 minutes

`$ docker run -d de13/spaceoddity:v0.1 -cpu -s 600` # run a memory leak app during 10 minutes