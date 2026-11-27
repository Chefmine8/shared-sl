#! /bin/sh

curl -O http://$1/sl
curl -O http://$1/sl-net
chmod +x sl
chmod +x sl-net

./sl-net join $1 0
