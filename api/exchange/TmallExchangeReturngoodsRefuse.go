package exchange

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/exchange"
)

// Tmallexchangereturngoodsrefuse 卖家拒绝确认收货
// tmall.exchange.returngoods.refuse
//
// 卖家拒绝买家换货申请
func Tmallexchangereturngoodsrefuse(clt *core.SDKClient, req *exchange.TmallexchangereturngoodsrefuseAPIRequest, session string) (*exchange.TmallexchangereturngoodsrefuseAPIResponse, error) {
	var resp exchange.TmallexchangereturngoodsrefuseAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
