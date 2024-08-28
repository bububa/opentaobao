package wdk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaRetailMarketingBuygiftActivitySkuAdd 添加单品买赠活动商品
// alibaba.retail.marketing.buygift.activity.sku.add
//
// 新增或更新单品买赠活动商品信息【同城零售】
func AlibabaRetailMarketingBuygiftActivitySkuAdd(ctx context.Context, clt *core.SDKClient, req *wdk.AlibabaRetailMarketingBuygiftActivitySkuAddAPIRequest, resp *wdk.AlibabaRetailMarketingBuygiftActivitySkuAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
