package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItemdiscountActivityQuery 查询单品特价活动【同城零售】
// alibaba.retail.marketing.itemdiscount.activity.query
//
// 查询单品特价活动【同城零售】
func AlibabaRetailMarketingItemdiscountActivityQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountActivityQueryAPIRequest, resp *wdk.AlibabaRetailMarketingItemdiscountActivityQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
