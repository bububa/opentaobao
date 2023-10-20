package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelOrderDetailSearch 订单详情查询
// taobao.xhotel.order.detail.search
//
// 提供订单详情查询
func TaobaoXhotelOrderDetailSearch(clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderDetailSearchAPIRequest, session string) (*xhotelonlineorder.TaobaoXhotelOrderDetailSearchAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoXhotelOrderDetailSearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
