package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftQueryitems 查询买赠活动下的商品
// alibaba.hm.marketing.itembuygift.queryitems
//
// 查询买赠活动下的商品
func AlibabaHmMarketingItembuygiftQueryitems(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftQueryitemsAPIRequest, resp *wdk.AlibabaHmMarketingItembuygiftQueryitemsAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
