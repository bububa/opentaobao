package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderAggregate 淘鲜达订单按门店机台号聚合查询
// alibaba.wdk.order.aggregate
//
// 淘鲜达订单按门店机台号聚合查询
func AlibabaWdkOrderAggregate(clt *core.SDKClient, req *wdk.AlibabaWdkOrderAggregateAPIRequest, resp *wdk.AlibabaWdkOrderAggregateAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
