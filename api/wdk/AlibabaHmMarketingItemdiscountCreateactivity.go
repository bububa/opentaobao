package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountCreateactivity 创建商品特价活动
// alibaba.hm.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
func AlibabaHmMarketingItemdiscountCreateactivity(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountCreateactivityAPIRequest, resp *wdk.AlibabaHmMarketingItemdiscountCreateactivityAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
