package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// TmallExchangeConfirmConsign 换货商家确认收货并发货
// tmall.exchange.confirm.consign
//
// 卖家确认收货并发货
func TmallExchangeConfirmConsign(clt *core.SDKClient, req *exchange.TmallExchangeConfirmConsignAPIRequest, session string) (*exchange.TmallExchangeConfirmConsignAPIResponse, error) {
	var resp exchange.TmallExchangeConfirmConsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
