package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderSync 五道口外部订单同步
// alibaba.wdk.order.sync
//
// 外部商户使用自助POS下单订单同步到五道口
func AlibabaWdkOrderSync(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkOrderSyncAPIRequest, resp *wdk.AlibabaWdkOrderSyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
