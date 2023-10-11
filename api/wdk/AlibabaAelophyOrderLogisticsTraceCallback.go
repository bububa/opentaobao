package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaAelophyOrderLogisticsTraceCallback 配送轨迹回传
// alibaba.aelophy.order.logistics.trace.callback
//
// 配送轨迹回传
func AlibabaAelophyOrderLogisticsTraceCallback(clt *core.SDKClient, req *wdk.AlibabaAelophyOrderLogisticsTraceCallbackAPIRequest, session string) (*wdk.AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse, error) {
	var resp wdk.AlibabaAelophyOrderLogisticsTraceCallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
