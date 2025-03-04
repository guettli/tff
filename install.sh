#!/usr/bin/env bash
# Bash Strict Mode: https://github.com/guettli/bash-strict-mode
trap 'echo "Warning: A command has failed. Exiting the script. Line was ($0:$LINENO): $(sed -n "${LINENO}p" "$0")"; exit 3' ERR
set -Eeuo pipefail

./test.sh

if [[ -n $(git status --porcelain) ]]; then
    echo "Error: Git repository is not clean. Please commit or stash your changes."
    exit 1
fi

go install .

sudo cp ten-flying-fingers.service /etc/systemd/system/
sudo systemctl daemon-reload
sudo systemctl enable ten-flying-fingers
sudo systemctl restart ten-flying-fingers
sudo systemctl status ten-flying-fingers --no-pager
