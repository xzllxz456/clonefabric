#!/bin/bash
set -ev

docker-compose -f docker-compose-ca.yaml up -d ca.sales1.acornpub.com

sleep 1
cd /home/rathon/go/src/github.com/hyperledger-fabric-stream-music/stream-music-2/application/sdk
node enrollAdmin.js
sleep 1
node registUsers.js
