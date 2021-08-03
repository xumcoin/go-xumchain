// Copyright 2015 The go-ethereum Authors
// This file is part of the go-ethereum library.
//
// The go-ethereum library is free software: you can redistribute it and/or modify
// it under the terms of the GNU Lesser General Public License as published by
// the Free Software Foundation, either version 3 of the License, or
// (at your option) any later version.
//
// The go-ethereum library is distributed in the hope that it will be useful,
// but WITHOUT ANY WARRANTY; without even the implied warranty of
// MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the
// GNU Lesser General Public License for more details.
//
// You should have received a copy of the GNU Lesser General Public License
// along with the go-ethereum library. If not, see <http://www.gnu.org/licenses/>.

package params

import "github.com/xumchain/go-xumchain/common"

// MainnetBootnodes are the enode URLs of the P2P bootstrap nodes running on
// the main Xumchain network.
var MainnetBootnodes = []string{
	// XUM Foundation Go Bootnodes
	"enode://0b0df8866fdaa73076e05b4b62f061642a501dda800c04e0b4ce6f993f29f653ce8b9ca972a97f39c4e5d6059b0762569687d74ca54c5e70f7737a39c4187b23@3.133.37.147:10101",   // bootnode-node1
	"enode://321103881513ce7c4d2bc053a7eac96833968dad4f38ff120d2e52425389db2c75791a8dafa7fba420d48a1665c63dc115c605a5cfbd2adfe80cfe77b2e0b11e@18.221.29.97:10101",     // bootnode-node2
}

// RopstenBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Ropsten test network.
var RopstenBootnodes = []string{
}

// RinkebyBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Rinkeby test network.
var RinkebyBootnodes = []string{
"enode://56312e7cef26a7945ee549edbe957e4f6655d52f4d4dfc3294c140dc1fb248a81a5deb6f4e111b1e9160611bc4d74dbc3ca896a86d5bff2d08bd05366a88cfb8@18.222.25.39:30303",
"enode://ec68ca788b7a56d923d6537bdaee861e6ca1a37aaeba8fba7b82b31c58c488b5da9e6a8136f8d69746f0785a072016cc661e6e0b412b9038a8981814bd4cbf3c@3.141.43.85:30303",
}

// GoerliBootnodes are the enode URLs of the P2P bootstrap nodes running on the
// Görli test network.
var GoerliBootnodes = []string{
	// Upstream bootnodes

	// Xumchain Foundation bootnode

	// Goerli Initiative bootnodes
}

// YoloV3Bootnodes are the enode URLs of the P2P bootstrap nodes running on the
// YOLOv3 ephemeral test network.
// TODO: Set Yolov3 bootnodes
var YoloV3Bootnodes = []string{
	"enode://9e1096aa59862a6f164994cb5cb16f5124d6c992cdbf4535ff7dea43ea1512afe5448dca9df1b7ab0726129603f1a3336b631e4d7a1a44c94daddd03241587f9@3.9.20.133:30303",
}

var V5Bootnodes = []string{
	// Teku team's bootnode
}

const dnsPrefix = "enrtree://AKA3AM6LPBYEUDMVNU3BSVQJ5AD45Y7YPOHJLEF6W26QOE4VTUDPE@"

// KnownDNSNetwork returns the address of a public DNS-based node list for the given
// genesis hash and protocol. See https://github.com/ethereum/discv4-dns-lists for more
// information.
func KnownDNSNetwork(genesis common.Hash, protocol string) string {
	var net string
	switch genesis {
	case MainnetGenesisHash:
		net = "mainnet"
	case RopstenGenesisHash:
		net = "ropsten"
	case RinkebyGenesisHash:
		net = "rinkeby"
	case GoerliGenesisHash:
		net = "goerli"
	default:
		return ""
	}
	return dnsPrefix + protocol + "." + net + ".ethdisco.net"
}
