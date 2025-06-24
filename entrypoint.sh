#!/usr/bin/sh
set -e
if [ !  -e /home/callisto/.initialized ]; then
    callisto init --home /etc/callisto
    callisto parse genesis-file --home /etc/callisto --genesis-file-path /callisto/genesis.json
    touch /home/callisto/.initialized
fi
exec $@
