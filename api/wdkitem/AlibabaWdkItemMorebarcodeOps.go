package wdkitem

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdkitem"
)

// AlibabaWdkItemMorebarcodeOps 商品一品多码维护操作
// alibaba.wdk.item.morebarcode.ops
//
// 商品一品多码维护操作
func AlibabaWdkItemMorebarcodeOps(ctx context.Context, clt *core.SDKClient, req *wdkitem.AlibabaWdkItemMorebarcodeOpsAPIRequest, resp *wdkitem.AlibabaWdkItemMorebarcodeOpsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
