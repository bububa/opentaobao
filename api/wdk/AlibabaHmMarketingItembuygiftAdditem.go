package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftAdditem 增加买赠活动商品。【注意，此接口暂不支持并发！】
// alibaba.hm.marketing.itembuygift.additem
//
// 增加买赠活动商品。【注意，此接口暂不支持并发！】
func AlibabaHmMarketingItembuygiftAdditem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftAdditemAPIRequest, resp *wdk.AlibabaHmMarketingItembuygiftAdditemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
