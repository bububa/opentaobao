package xhotelonlineorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelOrderSearch 酒店产品库订单查询
// taobao.xhotel.order.search
//
// 酒店产品库订单查询功能，查询90天内的订单
func TaobaoXhotelOrderSearch(ctx context.Context, clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderSearchAPIRequest, resp *xhotelonlineorder.TaobaoXhotelOrderSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
