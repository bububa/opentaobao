package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAxWarehouseOutboundCallback 翱象出仓回传
// alibaba.ax.warehouse.outbound.callback
//
// 翱象出仓回传
func AlibabaAxWarehouseOutboundCallback(clt *core.SDKClient, req *wdk.AlibabaAxWarehouseOutboundCallbackAPIRequest, resp *wdk.AlibabaAxWarehouseOutboundCallbackAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
