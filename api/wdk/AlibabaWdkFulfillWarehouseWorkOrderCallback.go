package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfillwarehouseworkordercallback 标准化仓作业单回传接口
// alibaba.wdk.fulfill.warehouse.work.order.callback
//
// 标准化仓作业单回传接口
func Alibabawdkfulfillwarehouseworkordercallback(clt *core.SDKClient, req *wdk.AlibabawdkfulfillwarehouseworkordercallbackAPIRequest, session string) (*wdk.AlibabawdkfulfillwarehouseworkordercallbackAPIResponse, error) {
	var resp wdk.AlibabawdkfulfillwarehouseworkordercallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
