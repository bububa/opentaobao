package wms

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wms"
)

// TaobaoWlbWmsOrderCancelNotify 单据取消接口
// taobao.wlb.wms.order.cancel.notify
//
// 单据取消接口
func TaobaoWlbWmsOrderCancelNotify(clt *core.SDKClient, req *wms.TaobaoWlbWmsOrderCancelNotifyAPIRequest, resp *wms.TaobaoWlbWmsOrderCancelNotifyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
