package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingBuygiftItemRemoveAsync 批量删除买赠商品
// alibaba.wdk.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
func AlibabaWdkMarketingBuygiftItemRemoveAsync(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIRequest, resp *wdk.AlibabaWdkMarketingBuygiftItemRemoveAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
