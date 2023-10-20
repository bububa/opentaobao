package axintrade

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/axintrade"
)

// Taobaoalitriptravelaxinhotelorderdetail 阿信酒店分销-订单详情接口
// taobao.alitrip.travel.axin.hotel.order.detail
//
// 阿信酒店订单详情的读取接口
func Taobaoalitriptravelaxinhotelorderdetail(clt *core.SDKClient, req *axintrade.TaobaoalitriptravelaxinhotelorderdetailAPIRequest, session string) (*axintrade.TaobaoalitriptravelaxinhotelorderdetailAPIResponse, error) {
	var resp axintrade.TaobaoalitriptravelaxinhotelorderdetailAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
