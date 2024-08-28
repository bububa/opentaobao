package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItemdiscountActivityUpdate 更新单品特价活动【同城零售】
// alibaba.retail.marketing.itemdiscount.activity.update
//
// 同城零售单品特价活动更新
func AlibabaRetailMarketingItemdiscountActivityUpdate(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountActivityUpdateAPIRequest, resp *wdk.AlibabaRetailMarketingItemdiscountActivityUpdateAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
