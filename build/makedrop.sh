#!/bin/bash

echo "wrapper script around make and docker issues"

DOCKER=`which docker`
MAKE=`which make`


echo " make found in "+$MAKE

echo " docker found in "+$DOCKER


exit