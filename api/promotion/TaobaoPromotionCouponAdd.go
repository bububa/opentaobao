package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TaobaoPromotionCouponAdd 创建店铺优惠券接口
// taobao.promotion.coupon.add
//
// 创建店铺优惠券。有效期内的店铺优惠券总数量不超过50张
func TaobaoPromotionCouponAdd(ctx context.Context, clt *core.SDKClient, req *promotion.TaobaoPromotionCouponAddAPIRequest, resp *promotion.TaobaoPromotionCouponAddAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
