package ton

import (
	"context"
	"fmt"

	"www.miniton-gateway.com/pkg/log"

	"github.com/xssnick/tonutils-go/liteclient"
	"github.com/xssnick/tonutils-go/tlb"
	"github.com/xssnick/tonutils-go/ton"
	"www.miniton-gateway.com/pkg/config"
)

var (
	block *tlb.BlockInfo
	ctx   context.Context
	api   ton.APIClientWrapped
)

func Init() {
	tonConfig := config.Config.TonConfig
	client := liteclient.NewConnectionPool()
	cfg, err := liteclient.GetConfigFromUrl(context.Background(), tonConfig.URL)
	if err != nil {
		panic(fmt.Sprintf("ton GetConfigFromUrl err,err is %v", err.Error()))
	}
	err = client.AddConnectionsFromConfig(context.Background(), cfg)
	if err != nil {
		panic(fmt.Sprintf("ton AddConnectionsFromConfig err,err is %v", err.Error()))
	}

	// api client with full proof checks
	api = ton.NewAPIClient(client, ton.ProofCheckPolicySecure).WithRetry()
	api.SetTrustedBlockFromConfig(cfg)

	// bound all requests to single ton node
	ctx = client.StickyContext(context.Background())

	log.Info(context.Background(), "ton start CurrentMasterchainInfo")

	block, err = api.CurrentMasterchainInfo(context.Background())
	if err != nil {
		panic(fmt.Sprintf("ton CurrentMasterchainInfo err,err is %v", err.Error()))
	}
	log.Info(context.Background(), "ton CurrentMasterchainInfo done")
}
