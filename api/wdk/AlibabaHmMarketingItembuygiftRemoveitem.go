package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftRemoveitem 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
// alibaba.hm.marketing.itembuygift.removeitem
//
// 移除买赠活动下的商品。【注意，此接口暂不支持并发！】
func AlibabaHmMarketingItembuygiftRemoveitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftRemoveitemAPIRequest, resp *wdk.AlibabaHmMarketingItembuygiftRemoveitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
