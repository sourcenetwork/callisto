#!/usr/bin/sh
set -e
if [ !  -e /etc/callisto/config.yaml ]; then
    callisto init --home /etc/callisto
fi

if [ !  -e /home/callisto/.initialized ]; then
    callisto parse genesis-file --home /etc/callisto --genesis-file-path /callisto/genesis.json
    touch /home/callisto/.initialized
fi
exec $@
