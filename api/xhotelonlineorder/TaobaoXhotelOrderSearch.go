package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelordersearch 酒店产品库订单查询
// taobao.xhotel.order.search
//
// 酒店产品库订单查询功能，查询90天内的订单
func Taobaoxhotelordersearch(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelordersearchAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelordersearchAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelordersearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
