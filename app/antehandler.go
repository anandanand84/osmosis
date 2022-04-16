package app

import (
	"log"
	"net/http"

	sdk "github.com/cosmos/cosmos-sdk/types"
)

type DvNotifyDecorator struct {
	url string
}

func NewDvNotifyDecorator(url string) DvNotifyDecorator {
	return DvNotifyDecorator{
		url: url,
	}
}

func (dvDec DvNotifyDecorator) AnteHandle(ctx sdk.Context, tx sdk.Tx, simulate bool, next sdk.AnteHandler) (sdk.Context, error) {
	log.Println("Informing dv endpoint of mempool transaction")
	_, err := http.Get(dvDec.url)
	if err != nil {
		log.Fatalln(err)
	}
	return next(ctx, tx, simulate)
}
