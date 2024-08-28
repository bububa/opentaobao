package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingBuygiftItemRemoveAsync 批量删除买赠商品
// alibaba.hm.marketing.buygift.item.remove.async
//
// 批量删除买赠商品
func AlibabaHmMarketingBuygiftItemRemoveAsync(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingBuygiftItemRemoveAsyncAPIRequest, resp *wdk.AlibabaHmMarketingBuygiftItemRemoveAsyncAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
