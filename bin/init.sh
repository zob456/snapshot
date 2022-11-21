#bin/bash

set -e

echo "Welcome to the snapshot project software validator!! Checking your installed software now..."
sleep 1;

if [[ $(which docker) ]]; then
  echo "Docker installed installed"!
else
  echo "Docker desktop NOT FOUND. It is required to run this project. You can download it at https://www.docker.com/products/docker-desktop/"
  exit 1;
fi

if [[ $(which tilt) ]]; then
  echo "tilt installed"!
else
  echo "Installing tilt & running command 'kubectl config use-context docker-desktop;'..."
  kubectl config use-context docker-desktop;
  curl -fsSL https://raw.githubusercontent.com/tilt-dev/tilt/master/scripts/install.sh | bash
fi
