package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkrefundaggregate 淘鲜达退款单按门店聚合查询
// alibaba.wdk.refund.aggregate
//
// 淘鲜达退款单按门店聚合查询
func Alibabawdkrefundaggregate(clt *core.SDKClient, req *wdk.AlibabawdkrefundaggregateAPIRequest, session string) (*wdk.AlibabawdkrefundaggregateAPIResponse, error) {
	var resp wdk.AlibabawdkrefundaggregateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
