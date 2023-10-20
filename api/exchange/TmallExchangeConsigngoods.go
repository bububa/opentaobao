package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangeconsigngoods 卖家发货
// tmall.exchange.consigngoods
//
// 卖家发货
func Tmallexchangeconsigngoods(clt *core.SDKClient, req *exchange.TmallexchangeconsigngoodsAPIRequest, session string) (*exchange.TmallexchangeconsigngoodsAPIResponse, error) {
	var resp exchange.TmallexchangeconsigngoodsAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
