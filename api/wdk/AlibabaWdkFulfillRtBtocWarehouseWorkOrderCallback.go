package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfillrtbtocwarehouseworkordercallback 大润发B2C仓作业单回传接口
// alibaba.wdk.fulfill.rt.btoc.warehouse.work.order.callback
//
// 大润发B2C仓作业单回传接口
func Alibabawdkfulfillrtbtocwarehouseworkordercallback(clt *core.SDKClient, req *wdk.AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIRequest, session string) (*wdk.AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIResponse, error) {
	var resp wdk.AlibabawdkfulfillrtbtocwarehouseworkordercallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
