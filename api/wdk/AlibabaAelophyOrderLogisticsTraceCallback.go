package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabaaelophyorderlogisticstracecallback 配送轨迹回传
// alibaba.aelophy.order.logistics.trace.callback
//
// 配送轨迹回传
func Alibabaaelophyorderlogisticstracecallback(clt *core.SDKClient, req *wdk.AlibabaaelophyorderlogisticstracecallbackAPIRequest, session string) (*wdk.AlibabaaelophyorderlogisticstracecallbackAPIResponse, error) {
	var resp wdk.AlibabaaelophyorderlogisticstracecallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
