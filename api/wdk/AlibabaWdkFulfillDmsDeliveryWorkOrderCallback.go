package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfilldmsdeliveryworkordercallback 末端配配送作业回传
// alibaba.wdk.fulfill.dms.delivery.work.order.callback
//
// 末端配配送作业回传。
func Alibabawdkfulfilldmsdeliveryworkordercallback(clt *core.SDKClient, req *wdk.AlibabawdkfulfilldmsdeliveryworkordercallbackAPIRequest, session string) (*wdk.AlibabawdkfulfilldmsdeliveryworkordercallbackAPIResponse, error) {
	var resp wdk.AlibabawdkfulfilldmsdeliveryworkordercallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
