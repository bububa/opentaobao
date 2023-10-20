package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkorderaggregate 淘鲜达订单按门店机台号聚合查询
// alibaba.wdk.order.aggregate
//
// 淘鲜达订单按门店机台号聚合查询
func Alibabawdkorderaggregate(clt *core.SDKClient, req *wdk.AlibabawdkorderaggregateAPIRequest, session string) (*wdk.AlibabawdkorderaggregateAPIResponse, error) {
	var resp wdk.AlibabawdkorderaggregateAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
