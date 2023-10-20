package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitriptravelaxinhotelticketrefundorderrefund 阿信度假业务申请退款
// taobao.alitrip.travel.axin.hotelticket.refund.orderrefund
//
// 阿信度假业务申请退款
func Taobaoalitriptravelaxinhotelticketrefundorderrefund(clt *core.SDKClient, req *axintrade.TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIRequest, session string) (*axintrade.TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIResponse, error) {
	var resp axintrade.TaobaoalitriptravelaxinhotelticketrefundorderrefundAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
