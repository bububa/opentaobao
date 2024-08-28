package wms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsStockOutOrderNotify 出库单通知
// taobao.wlb.wms.stock.out.order.notify
//
// 出库单通知
func TaobaoWlbWmsStockOutOrderNotify(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsStockOutOrderNotifyAPIRequest, resp *wms.TaobaoWlbWmsStockOutOrderNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
