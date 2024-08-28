package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOrderSyncWithitem 订单和商品同步接口
// alibaba.wdk.order.sync.withitem
//
// 轻pos,将订单和商品的信息一起传到盒马这边，进行创单和添加商品处理。
func AlibabaWdkOrderSyncWithitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkOrderSyncWithitemAPIRequest, resp *wdk.AlibabaWdkOrderSyncWithitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
