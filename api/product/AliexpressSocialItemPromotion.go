package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// AliexpressSocialItemPromotion 获取推广链接
// aliexpress.social.item.promotion
//
// 获取商品社交推广链接
func AliexpressSocialItemPromotion(ctx context.Context, clt *core.SDKClient, req *product.AliexpressSocialItemPromotionAPIRequest, resp *product.AliexpressSocialItemPromotionAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
