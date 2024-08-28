package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderRefundList 五道口交易退款批量查询
// alibaba.wdk.order.refund.list
//
// 按照条件查询退款数据
func AlibabaWdkOrderRefundList(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkOrderRefundListAPIRequest, resp *wdk.AlibabaWdkOrderRefundListAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
