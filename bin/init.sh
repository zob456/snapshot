#bin/bash

set -e

kubectl config use-context docker-desktop;

curl -fsSL https://raw.githubusercontent.com/tilt-dev/tilt/master/scripts/install.sh | bash