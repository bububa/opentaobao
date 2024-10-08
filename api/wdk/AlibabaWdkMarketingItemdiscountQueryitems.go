package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItemdiscountQueryitems 查询特价商品
// alibaba.wdk.marketing.itemdiscount.queryitems
//
// 查询参加特价活动的商品优惠详情
func AlibabaWdkMarketingItemdiscountQueryitems(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountQueryitemsAPIRequest, resp *wdk.AlibabaWdkMarketingItemdiscountQueryitemsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
