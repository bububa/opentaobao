package product

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/product"
)

// TaobaoUmpPromotionSkuGet 商品优惠详情查询
// taobao.ump.promotion.sku.get
//
// 商品优惠详情查询，可查询商品设置的详细优惠。包括限时折扣，满就送等官方优惠以及第三方优惠。
func TaobaoUmpPromotionSkuGet(ctx context.Context, clt *core.SDKClient, req *product.TaobaoUmpPromotionSkuGetAPIRequest, resp *product.TaobaoUmpPromotionSkuGetAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
