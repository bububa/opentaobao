package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsReturnOrderNotify 销售退货通知
// taobao.wlb.wms.return.order.notify
//
// 销售退货通知
func TaobaoWlbWmsReturnOrderNotify(clt *core.SDKClient, req *wms.TaobaoWlbWmsReturnOrderNotifyAPIRequest, resp *wms.TaobaoWlbWmsReturnOrderNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
