package ton

import (
	"github.com/xssnick/tonutils-go/address"
	"github.com/xssnick/tonutils-go/ton/wallet"
)

func CreateAccount() (address string, words []string, err error) {
	words = wallet.NewSeed()
	w, err := wallet.FromSeed(api, words, wallet.V3)
	if err != nil {
		return
	}
	address = w.Address().String()
	return
}

func GetAccountBalance(userAddress string) (balance string, err error) {
	addr := address.MustParseAddr(userAddress)

	res, err := api.WaitForBlock(block.SeqNo).GetAccount(ctx, block, addr)
	if err != nil {
		return
	}
	if res.IsActive {
		balance = res.State.Balance.String()
	}
	return
}
