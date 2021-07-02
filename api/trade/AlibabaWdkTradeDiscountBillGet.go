package trade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/trade"
)

// AlibabaWdkTradeDiscountBillGet 订单优惠账单查询
// alibaba.wdk.trade.discount.bill.get
//
// 商家查询订单优惠账单
func AlibabaWdkTradeDiscountBillGet(clt *core.SDKClient, req *trade.AlibabaWdkTradeDiscountBillGetAPIRequest, session string) (*trade.AlibabaWdkTradeDiscountBillGetAPIResponse, error) {
	var resp trade.AlibabaWdkTradeDiscountBillGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
