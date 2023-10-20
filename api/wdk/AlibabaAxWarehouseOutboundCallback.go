package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaxwarehouseoutboundcallback 翱象出仓回传
// alibaba.ax.warehouse.outbound.callback
//
// 翱象出仓回传
func Alibabaaxwarehouseoutboundcallback(clt *core.SDKClient, req *wdk.AlibabaaxwarehouseoutboundcallbackAPIRequest, session string) (*wdk.AlibabaaxwarehouseoutboundcallbackAPIResponse, error) {
	var resp wdk.AlibabaaxwarehouseoutboundcallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
