#! /bin/sh

curl -O http://$1/nix/sl
curl -O http://$1/nix/sl-net
chmod +x sl
chmod +x sl-net

./sl-net join $1 0
