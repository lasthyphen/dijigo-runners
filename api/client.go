package api

import (
	"github.com/lasthyphen/beacongo/api/admin"
	"github.com/lasthyphen/beacongo/api/health"
	"github.com/lasthyphen/beacongo/api/info"
	"github.com/lasthyphen/beacongo/api/ipcs"
	"github.com/lasthyphen/beacongo/api/keystore"
	"github.com/lasthyphen/beacongo/indexer"
	"github.com/lasthyphen/beacongo/vms/avm"
	"github.com/lasthyphen/beacongo/vms/platformvm"
	"github.com/lasthyphen/coreth/plugin/evm"
)

// Issues API calls to a node
// TODO: byzantine api. check if appropiate. improve implementation.
type Client interface {
	PChainAPI() platformvm.Client
	XChainAPI() avm.Client
	XChainWalletAPI() avm.WalletClient
	CChainAPI() evm.Client
	CChainEthAPI() EthClient // ethclient websocket wrapper that adds mutexed calls, and lazy conn init (on first call)
	InfoAPI() info.Client
	HealthAPI() health.Client
	IpcsAPI() ipcs.Client
	KeystoreAPI() keystore.Client
	AdminAPI() admin.Client
	PChainIndexAPI() indexer.Client
	CChainIndexAPI() indexer.Client
	// TODO add methods
}
