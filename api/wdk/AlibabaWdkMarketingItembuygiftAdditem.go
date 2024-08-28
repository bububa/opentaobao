package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItembuygiftAdditem 增加买赠活动商品。【注意，此接口暂不支持并发！】
// alibaba.wdk.marketing.itembuygift.additem
//
// 增加买赠活动商品。【注意，此接口暂不支持并发！】
func AlibabaWdkMarketingItembuygiftAdditem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftAdditemAPIRequest, resp *wdk.AlibabaWdkMarketingItembuygiftAdditemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
