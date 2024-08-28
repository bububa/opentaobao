package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderAggregate 淘鲜达订单按门店机台号聚合查询
// alibaba.wdk.order.aggregate
//
// 淘鲜达订单按门店机台号聚合查询
func AlibabaWdkOrderAggregate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkOrderAggregateAPIRequest, resp *wdk.AlibabaWdkOrderAggregateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
