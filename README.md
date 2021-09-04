# XUMChain Mainnet - Integration guide


Connecting to the XUM network is extremely easy and only requires to download the Linux executable and run the following commands.

git clone https://github.com/xumcoin/xum-exec

./gxum --syncmode 'full' --networkid 10001

Now, you will have the full XUM node running in your computer.

Mainnet Block Explorer: https://xumexplorer.com

Testnet Block Explorer: https://testnet.xumexplorer.com

# How to develop DAPPs in XUMChain

XUMChain is an EVM based blockchain and DAPPs can be developed using Web3.js library.

First you need to get web3.js into your project. This can be done using the following methods:

Install web3.js using npm 

1. npm install web3

2. and then the following commands.. 
var Web3 = require('web3');
var web3 = new Web3('http://localhost:8545');
// or
var web3 = new Web3(new Web3.providers.HttpProvider('http://localhost:8545'));


The web3.js library is a collection of modules that contain functionality for the EVM ecosystem.

Please go through the following material to read about Web3 in detail.

https://www.npmjs.com/package/web3

https://web3js.readthedocs.io/en/v1.5.2/getting-started.html





