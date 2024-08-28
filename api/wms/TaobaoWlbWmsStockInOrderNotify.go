package wms

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsStockInOrderNotify 入库通知单
// taobao.wlb.wms.stock.in.order.notify
//
// 入库通知单
func TaobaoWlbWmsStockInOrderNotify(ctx context.Context, clt *core.SDKClient, req *wms.TaobaoWlbWmsStockInOrderNotifyAPIRequest, resp *wms.TaobaoWlbWmsStockInOrderNotifyAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
