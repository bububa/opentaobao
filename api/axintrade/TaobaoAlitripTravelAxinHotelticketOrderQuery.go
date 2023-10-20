package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitriptravelaxinhotelticketorderquery 阿信度假交易订单查询接口
// taobao.alitrip.travel.axin.hotelticket.order.query
//
// 阿信度假交易订单查询接口
func Taobaoalitriptravelaxinhotelticketorderquery(clt *core.SDKClient, req *axintrade.TaobaoalitriptravelaxinhotelticketorderqueryAPIRequest, session string) (*axintrade.TaobaoalitriptravelaxinhotelticketorderqueryAPIResponse, error) {
	var resp axintrade.TaobaoalitriptravelaxinhotelticketorderqueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
