package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkfulfillmissfreshwarehouseworkordercallback 每日优鲜仓作业单回传接口
// alibaba.wdk.fulfill.missfresh.warehouse.work.order.callback
//
// 家乐福仓作业单回传接口
func Alibabawdkfulfillmissfreshwarehouseworkordercallback(clt *core.SDKClient, req *wdk.AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIRequest, session string) (*wdk.AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIResponse, error) {
	var resp wdk.AlibabawdkfulfillmissfreshwarehouseworkordercallbackAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
