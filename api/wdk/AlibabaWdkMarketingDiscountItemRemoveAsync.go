package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingDiscountItemRemoveAsync 特价批量移除商品
// alibaba.wdk.marketing.discount.item.remove.async
//
// 特价批量移除商品
func AlibabaWdkMarketingDiscountItemRemoveAsync(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingDiscountItemRemoveAsyncAPIRequest, resp *wdk.AlibabaWdkMarketingDiscountItemRemoveAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
