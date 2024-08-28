package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItemdiscountActivityDelete 删除单品特价活动【同城零售】
// alibaba.retail.marketing.itemdiscount.activity.delete
//
// 同城零售单品特价活动删除
func AlibabaRetailMarketingItemdiscountActivityDelete(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountActivityDeleteAPIRequest, resp *wdk.AlibabaRetailMarketingItemdiscountActivityDeleteAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
