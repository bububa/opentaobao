package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsStockOutOrderNotify 出库单通知
// taobao.wlb.wms.stock.out.order.notify
//
// 出库单通知
func TaobaoWlbWmsStockOutOrderNotify(clt *core.SDKClient, req *wms.TaobaoWlbWmsStockOutOrderNotifyAPIRequest, resp *wms.TaobaoWlbWmsStockOutOrderNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
