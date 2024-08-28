package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountQueryactivity 查找特价活动
// alibaba.hm.marketing.itemdiscount.queryactivity
//
// 查找特价活动
func AlibabaHmMarketingItemdiscountQueryactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountQueryactivityAPIRequest, resp *wdk.AlibabaHmMarketingItemdiscountQueryactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
