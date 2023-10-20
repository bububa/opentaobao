package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangeconfirmconsign 换货商家确认收货并发货
// tmall.exchange.confirm.consign
//
// 卖家确认收货并发货
func Tmallexchangeconfirmconsign(clt *core.SDKClient, req *exchange.TmallexchangeconfirmconsignAPIRequest, session string) (*exchange.TmallexchangeconfirmconsignAPIResponse, error) {
	var resp exchange.TmallexchangeconfirmconsignAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
