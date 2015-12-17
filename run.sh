#!/bin/bash

docker run -d -v '/etc/ssl/certs/ca-certificates.crt:/etc/ssl/certs/ca-certificates.crt' -p 10000:10000 vnadgir/status-api
