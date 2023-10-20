package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitriptravelaxinhotelorderpay 阿信酒店分销订单支付
// taobao.alitrip.travel.axin.hotel.order.pay
//
// 阿信酒店分销订单支付
func Taobaoalitriptravelaxinhotelorderpay(clt *core.SDKClient, req *axintrade.TaobaoalitriptravelaxinhotelorderpayAPIRequest, session string) (*axintrade.TaobaoalitriptravelaxinhotelorderpayAPIResponse, error) {
	var resp axintrade.TaobaoalitriptravelaxinhotelorderpayAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
