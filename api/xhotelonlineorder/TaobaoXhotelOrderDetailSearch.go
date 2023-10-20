package xhotelonlineorder

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// Taobaoxhotelorderdetailsearch 订单详情查询
// taobao.xhotel.order.detail.search
//
// 提供订单详情查询
func Taobaoxhotelorderdetailsearch(clt *core.SDKClient, req *xhotelonlineorder.TaobaoxhotelorderdetailsearchAPIRequest, session string) (*xhotelonlineorder.TaobaoxhotelorderdetailsearchAPIResponse, error) {
	var resp xhotelonlineorder.TaobaoxhotelorderdetailsearchAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
