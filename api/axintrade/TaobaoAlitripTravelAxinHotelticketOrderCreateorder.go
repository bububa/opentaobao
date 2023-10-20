package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitriptravelaxinhotelticketordercreateorder 阿信度假业务创单并支付接口
// taobao.alitrip.travel.axin.hotelticket.order.createorder
//
// 阿信度假业务创单并支付接口
func Taobaoalitriptravelaxinhotelticketordercreateorder(clt *core.SDKClient, req *axintrade.TaobaoalitriptravelaxinhotelticketordercreateorderAPIRequest, session string) (*axintrade.TaobaoalitriptravelaxinhotelticketordercreateorderAPIResponse, error) {
	var resp axintrade.TaobaoalitriptravelaxinhotelticketordercreateorderAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
