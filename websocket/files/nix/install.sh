#!/bin/sh

curl -O "http://$1/nix/sl"
curl -O "http://$1/nix/sl-net"
curl -O "http://$1/oxidvideos-train-over-bridge-inside-ambiance-617983.mp3"
curl -O "http://$1/caboose_canine-steam-locomotive-fouled-5-chime-375193.mp3"
chmod +x sl sl-net

nix profile add nixpkgs#alsa-lib nixpkgs#patchelf

INTERP=$(find /nix/store -maxdepth 3 -name "ld-linux-x86-64.so.2" 2>/dev/null | head -n 1)
if [ -n "$INTERP" ]; then
  patchelf --set-interpreter "$INTERP" ./sl-net
fi

if [ ! -f /etc/asound.conf ]; then
  cat << 'EOF' > /etc/asound.conf
pcm.!default {
    type plug
    slave.pcm "hw:0,3"
}

ctl.!default {
    type hw
    card 0
}
EOF
fi

ALSA_LIB_DIR=$(dirname "$(find /nix/store -maxdepth 3 -name "libasound.so.2" 2>/dev/null | head -n 1)")
export LD_LIBRARY_PATH="${ALSA_LIB_DIR}:$HOME/.nix-profile/lib:${LD_LIBRARY_PATH}"

./sl-net join "$1" 0
