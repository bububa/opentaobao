package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// Alibabawdktradediscountbillget 订单优惠账单查询
// alibaba.wdk.trade.discount.bill.get
//
// 商家查询订单优惠账单
func Alibabawdktradediscountbillget(clt *core.SDKClient, req *trade.AlibabawdktradediscountbillgetAPIRequest, session string) (*trade.AlibabawdktradediscountbillgetAPIResponse, error) {
	var resp trade.AlibabawdktradediscountbillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
