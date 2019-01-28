# Peer-to-peer file sharing system

P2P Sharer is a peer to peer file sharing system that is built on top of IPFS protocol which allows you to directly send your local file to other machine without having to rely on any third party organization like Dropbox and so on.

More than that, it also enables you to add your friend by linking your peers' ipfs ID with the naming you set.

### Prerequisites
You are required to download and install [ipfs](https://docs.ipfs.io/introduction/install/#installing-from-a-prebuilt-package), and host your own node by running ***ipfs daemon*** before running any command from p2p-sharer.

### Usage
Add file:
>p2p-sharer up "***file/directory name***"

Add friend:
>p2p-sharer addfriend "***peer's ID***" "***friend's name***"

Retrieve Object from your added friends:
>p2p-sharer retobject "***ipfs hash***" -n "***friend's name***" -o "***output file name***"

Get your own ipfs ID:
>p2p-sharer myid


## Main Motivation

This project is aimed to serve [Youdeez](https://github.com/moon004/YouDeez) as a medium for its users to share their music Playlists (including the song binaries) via peer to peer protocols and "friending" system.
