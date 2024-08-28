package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AliexpressSocialItemSearch AE社交选品
// aliexpress.social.item.search
//
// AE社交选品,通过各种筛选条件对社交商品池进行筛选
func AliexpressSocialItemSearch(ctx context.Context, clt *core.SDKClient, req *product.AliexpressSocialItemSearchAPIRequest, resp *product.AliexpressSocialItemSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
