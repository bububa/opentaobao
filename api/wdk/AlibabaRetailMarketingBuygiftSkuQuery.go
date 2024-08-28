package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftSkuQuery 查询买赠活动商品【同城零售】
// alibaba.retail.marketing.buygift.sku.query
//
// 查询买赠活动商品【同城零售】
func AlibabaRetailMarketingBuygiftSkuQuery(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftSkuQueryAPIRequest, resp *wdk.AlibabaRetailMarketingBuygiftSkuQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
