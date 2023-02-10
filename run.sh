#!/bin/sh

docker build -t replicate . && docker run -v `pwd`/tests/:/tests -p 9191:8080 replicate