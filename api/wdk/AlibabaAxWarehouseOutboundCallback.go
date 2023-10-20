package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAxWarehouseOutboundCallback 翱象出仓回传
// alibaba.ax.warehouse.outbound.callback
//
// 翱象出仓回传
func AlibabaAxWarehouseOutboundCallback(clt *core.SDKClient, req *wdk.AlibabaAxWarehouseOutboundCallbackAPIRequest, session string) (*wdk.AlibabaAxWarehouseOutboundCallbackAPIResponse, error) {
	var resp wdk.AlibabaAxWarehouseOutboundCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
