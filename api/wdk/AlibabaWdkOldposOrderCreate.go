package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOldposOrderCreate 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
// alibaba.wdk.oldpos.order.create
//
// 淘鲜达外部商户老pos机产生的订单同步进淘鲜达
func AlibabaWdkOldposOrderCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkOldposOrderCreateAPIRequest, resp *wdk.AlibabaWdkOldposOrderCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
