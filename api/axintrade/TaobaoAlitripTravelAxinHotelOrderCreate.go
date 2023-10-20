package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitriptravelaxinhotelordercreate 酒店分销订单创建服务-阿信
// taobao.alitrip.travel.axin.hotel.order.create
//
// 提供酒店分销订单创建服务
func Taobaoalitriptravelaxinhotelordercreate(clt *core.SDKClient, req *axintrade.TaobaoalitriptravelaxinhotelordercreateAPIRequest, session string) (*axintrade.TaobaoalitriptravelaxinhotelordercreateAPIResponse, error) {
	var resp axintrade.TaobaoalitriptravelaxinhotelordercreateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
