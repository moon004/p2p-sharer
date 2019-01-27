# Peer-to-peer file sharing system

P2P Sharer is a peer to peer file sharing system that is built on top of IPFS protocol which allows you to directly send your local file to other machine without having to rely on any third party organization like Dropbox and so on.

More than that, it also enables you to add your friend by linking your peers' ipfs ID with the naming you set.

### Prerequisites
You are required to download and install [ipfs](https://docs.ipfs.io/introduction/install/#installing-from-a-prebuilt-package), and host your own node by running ***ipfs daemon*** before running any command from p2p-sharer.

### Milestone

1. upfile (Finished)

Add the file to your local node and publish it to the network so that your peers are able to retrieve it.

2. retobject (Finished not yet tested) a.k.a ***retrieve object***

Connect to your ipfs peers (added friends) and get the file from the hash.

3. friend System(Finished)

Add your peers as friends and refer them based on your naming.


## Main Motivation

This project is aimed to serve [Youdeez](https://github.com/moon004/YouDeez) as a medium for its users to share their music Playlists (including the song binaries) via peer to peer protocols and "friending" system.
