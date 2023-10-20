package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitriptravelaxinhotelticketproductdetail 阿信酒景套餐产品详情查询
// taobao.alitrip.travel.axin.hotelticket.product.detail
//
// 阿信酒景套餐产品详情查询
func Taobaoalitriptravelaxinhotelticketproductdetail(clt *core.SDKClient, req *axintrade.TaobaoalitriptravelaxinhotelticketproductdetailAPIRequest, session string) (*axintrade.TaobaoalitriptravelaxinhotelticketproductdetailAPIResponse, error) {
	var resp axintrade.TaobaoalitriptravelaxinhotelticketproductdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
