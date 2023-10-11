package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAxWarehouseInboundCallback 翱象入库回传
// alibaba.ax.warehouse.inbound.callback
//
// 翱象入库回传
func AlibabaAxWarehouseInboundCallback(clt *core.SDKClient, req *wdk.AlibabaAxWarehouseInboundCallbackAPIRequest, session string) (*wdk.AlibabaAxWarehouseInboundCallbackAPIResponse, error) {
	var resp wdk.AlibabaAxWarehouseInboundCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
