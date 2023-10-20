package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// Taobaowlbwmsstockoutbillget 通过订单号获取单个出库单发货信息
// taobao.wlb.wms.stock.out.bill.get
//
// 通过订单号获取单个出库单发货信息
func Taobaowlbwmsstockoutbillget(clt *core.SDKClient, req *wms.TaobaowlbwmsstockoutbillgetAPIRequest, session string) (*wms.TaobaowlbwmsstockoutbillgetAPIResponse, error) {
	var resp wms.TaobaowlbwmsstockoutbillgetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
