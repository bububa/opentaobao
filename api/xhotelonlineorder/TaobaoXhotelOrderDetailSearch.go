package xhotelonlineorder

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/xhotelonlineorder"
)

// TaobaoXhotelOrderDetailSearch 订单详情查询
// taobao.xhotel.order.detail.search
//
// 提供订单详情查询
func TaobaoXhotelOrderDetailSearch(ctx context.Context, clt *core.SDKClient, req *xhotelonlineorder.TaobaoXhotelOrderDetailSearchAPIRequest, resp *xhotelonlineorder.TaobaoXhotelOrderDetailSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
