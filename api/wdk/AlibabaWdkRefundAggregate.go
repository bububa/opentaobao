package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkRefundAggregate 淘鲜达退款单按门店聚合查询
// alibaba.wdk.refund.aggregate
//
// 淘鲜达退款单按门店聚合查询
func AlibabaWdkRefundAggregate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkRefundAggregateAPIRequest, resp *wdk.AlibabaWdkRefundAggregateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
