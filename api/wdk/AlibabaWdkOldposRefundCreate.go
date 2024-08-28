package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkOldposRefundCreate 五道口外部商户老pos机产生的退款单同步进盒马
// alibaba.wdk.oldpos.refund.create
//
// 淘鲜达外部商户老pos机产生的退款单同步进淘鲜达
func AlibabaWdkOldposRefundCreate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkOldposRefundCreateAPIRequest, resp *wdk.AlibabaWdkOldposRefundCreateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
