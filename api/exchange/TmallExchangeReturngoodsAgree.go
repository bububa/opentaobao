package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangereturngoodsagree 卖家确认收货
// tmall.exchange.returngoods.agree
//
// 卖家确认收货
func Tmallexchangereturngoodsagree(clt *core.SDKClient, req *exchange.TmallexchangereturngoodsagreeAPIRequest, session string) (*exchange.TmallexchangereturngoodsagreeAPIResponse, error) {
	var resp exchange.TmallexchangereturngoodsagreeAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
