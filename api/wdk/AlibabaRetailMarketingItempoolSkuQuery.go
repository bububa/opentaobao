package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingItempoolSkuQuery 查询商品池活动商品【同城零售】
// alibaba.retail.marketing.itempool.sku.query
//
// 查询商品池活动商品【同城零售】
func AlibabaRetailMarketingItempoolSkuQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingItempoolSkuQueryAPIRequest, resp *wdk.AlibabaRetailMarketingItempoolSkuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
