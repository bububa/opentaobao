package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionCouponSellerSearch 查询绑定卖家优惠券相关信息
// taobao.promotion.coupon.seller.search
//
// 查询绑定卖家相关优惠券信息  如isv  百川 等外部业务方
func TaobaoPromotionCouponSellerSearch(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionCouponSellerSearchAPIRequest, resp *promotion.TaobaoPromotionCouponSellerSearchAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
