package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderGet 交易订单详情查询
// alibaba.wdk.order.get
//
// 五道口三江单据查询接口
func AlibabaWdkOrderGet(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkOrderGetAPIRequest, resp *wdk.AlibabaWdkOrderGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
