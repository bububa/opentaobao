package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitriptravelaxinhotelticketordervalidate 阿信度假业务交易试单接口
// taobao.alitrip.travel.axin.hotelticket.order.validate
//
// 阿信度假业务交易试单接口
func Taobaoalitriptravelaxinhotelticketordervalidate(clt *core.SDKClient, req *axintrade.TaobaoalitriptravelaxinhotelticketordervalidateAPIRequest, session string) (*axintrade.TaobaoalitriptravelaxinhotelticketordervalidateAPIResponse, error) {
	var resp axintrade.TaobaoalitriptravelaxinhotelticketordervalidateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
