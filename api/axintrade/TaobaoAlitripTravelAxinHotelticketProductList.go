package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitriptravelaxinhotelticketproductlist 阿信酒景套餐产品列表查询
// taobao.alitrip.travel.axin.hotelticket.product.list
//
// 阿信酒景套餐产品列表查询
func Taobaoalitriptravelaxinhotelticketproductlist(clt *core.SDKClient, req *axintrade.TaobaoalitriptravelaxinhotelticketproductlistAPIRequest, session string) (*axintrade.TaobaoalitriptravelaxinhotelticketproductlistAPIResponse, error) {
	var resp axintrade.TaobaoalitriptravelaxinhotelticketproductlistAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
