package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountRemoveitem 移除报名的商品
// alibaba.hm.marketing.itemdiscount.removeitem
//
// 将报名特价活动的商品从特价活动中移除
func AlibabaHmMarketingItemdiscountRemoveitem(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountRemoveitemAPIRequest, resp *wdk.AlibabaHmMarketingItemdiscountRemoveitemAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
