package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItemdiscountQueryactivity 查找特价活动
// alibaba.wdk.marketing.itemdiscount.queryactivity
//
// 查找特价活动
func AlibabaWdkMarketingItemdiscountQueryactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountQueryactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItemdiscountQueryactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
