package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfilldmsebeecakeworkordercallback 北京小蜜蜂配作业回传
// alibaba.wdk.fulfill.dms.ebeecake.work.order.callback
//
// 北京小蜜蜂配作业回传。
func Alibabawdkfulfilldmsebeecakeworkordercallback(clt *core.SDKClient, req *wdk.AlibabawdkfulfilldmsebeecakeworkordercallbackAPIRequest, session string) (*wdk.AlibabawdkfulfilldmsebeecakeworkordercallbackAPIResponse, error) {
	var resp wdk.AlibabawdkfulfilldmsebeecakeworkordercallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
