package aecreatives

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/aecreatives"
)

// AliexpressAffiliateProductQuery 联盟推广商品获取接口
// aliexpress.affiliate.product.query
//
// 联盟推广商品搜索接口，用于搜索联盟推广商品数据
func AliexpressAffiliateProductQuery(ctx context.Context, clt *core.SDKClient, req *aecreatives.AliexpressAffiliateProductQueryAPIRequest, resp *aecreatives.AliexpressAffiliateProductQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
