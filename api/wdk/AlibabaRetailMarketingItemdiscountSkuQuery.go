package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItemdiscountSkuQuery 查询单品特价活动商品【同城零售】
// alibaba.retail.marketing.itemdiscount.sku.query
//
// 查询单品特价活动商品【同城零售】
func AlibabaRetailMarketingItemdiscountSkuQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItemdiscountSkuQueryAPIRequest, resp *wdk.AlibabaRetailMarketingItemdiscountSkuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
