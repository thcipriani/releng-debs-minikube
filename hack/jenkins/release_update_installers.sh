#!/bin/bash

# Copyright 2016 The Kubernetes Authors All rights reserved.
#
# Licensed under the Apache License, Version 2.0 (the "License");
# you may not use this file except in compliance with the License.
# You may obtain a copy of the License at
#
#     http://www.apache.org/licenses/LICENSE-2.0
#
# Unless required by applicable law or agreed to in writing, software
# distributed under the License is distributed on an "AS IS" BASIS,
# WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
# See the License for the specific language governing permissions and
# limitations under the License.

git config --global user.name "minikube-bot"
git config --global user.email "minikube-bot@google.com"

REPLACE_PKG_VERSION=${VERSION_MAJOR}.${VERSION_MINOR}.${VERSION_BUILD}
REPLACE_MINIKUBE_LINUX_SHA256=$(awk '{ print $1 }' out/minikube-linux-amd64.sha256)
REPLACE_MINIKUBE_DARWIN_SHA256=$(awk '{ print $1 }' out/minikube-darwin-amd64.sha256)
REPLACE_CASK_CHECKPOINT=$(curl \
                        --compressed \
                        --location   \
                        --user-agent 'Mozilla/5.0 (Macintosh; Intel Mac OS X 10_9_2) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/33.0.1750.152 Safari/537.36' \
                        https://github.com/kubernetes/minikube/releases.atom \
                        | sed 's|<pubDate>[^<]*</pubDate>||g' \
                        | shasum -a 256 | awk '{ print $1 }')
MINIKUBE_ROOT=$PWD

git clone ssh://aur@aur.archlinux.org/minikube.git aur-minikube
pushd aur-minikube >/dev/null
    sed -e "s/\$PKG_VERSION/${REPLACE_PKG_VERSION}/g" \
        -e "s/\$MINIKUBE_LINUX_SHA256/${REPLACE_MINIKUBE_LINUX_SHA256}/g" \
        $MINIKUBE_ROOT/installers/linux/archlinux/PKGBUILD > PKGBUILD
    sed -e "s/\$PKG_VERSION/${REPLACE_PKG_VERSION}/g" \
        -e "s/\$MINIKUBE_LINUX_SHA256/${REPLACE_MINIKUBE_LINUX_SHA256}/g" \
        $MINIKUBE_ROOT/installers/linux/archlinux/.SRCINFO > .SRCINFO
    git add PKGBUILD .SRCINFO
    git commit -m "Upgrade to version ${REPLACE_PKG_VERSION}"

    git push origin master

popd >/dev/null

git clone --depth 1 git@github.com:minikube-bot/homebrew-cask.git # dont't pull entire history

pushd homebrew-cask >/dev/null
    git remote add upstream https://github.com/caskroom/homebrew-cask.git
    git fetch upstream
    git checkout upstream/master
    git checkout -b ${REPLACE_PKG_VERSION}
    sed -e "s/\$PKG_VERSION/${REPLACE_PKG_VERSION}/g" \
        -e "s/\$MINIKUBE_DARWIN_SHA256/${REPLACE_MINIKUBE_DARWIN_SHA256}/g" \
        -e "s/\$CASK_CHECKPOINT/${REPLACE_CASK_CHECKPOINT}/g" \
        $MINIKUBE_ROOT/installers/darwin/brew-cask/minikube.rb.tmpl > Casks/minikube.rb
    git add Casks/minikube.rb
    git commit -F- <<EOF
Update minikube to ${REPLACE_PKG_VERSION}

- [x] brew cask audit --download {{cask_file}} is error-free.
- [x] brew cask style --fix {{cask_file}} reports no offenses.
- [x] The commit message includes the cask’s name and version.

EOF
    git push origin ${REPLACE_PKG_VERSION}
    curl -v -k -u minikube-bot:${BOT_PASSWORD} -X POST https://api.github.com/repos/caskroom/homebrew-cask/pulls \
    -d @- <<EOF

{
    "title": "Update minikube to ${REPLACE_PKG_VERSION}",
    "head": "minikube-bot:${REPLACE_PKG_VERSION}",
    "base": "master",
    "body": "cc @r2d4"
}

EOF
popd >/dev/null

rm -rf aur-minikube homebrew-cask
