package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitriptravelaxinhotelorderrefund 阿信酒店分销订单退款API
// taobao.alitrip.travel.axin.hotel.order.refund
//
// 阿信酒店分销订单退款
func Taobaoalitriptravelaxinhotelorderrefund(clt *core.SDKClient, req *axintrade.TaobaoalitriptravelaxinhotelorderrefundAPIRequest, session string) (*axintrade.TaobaoalitriptravelaxinhotelorderrefundAPIResponse, error) {
	var resp axintrade.TaobaoalitriptravelaxinhotelorderrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
