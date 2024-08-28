package promotion

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// TmallPromotionCouponQuery 查询可用优惠券列表
// tmall.promotion.coupon.query
//
// 查询用户的可用优惠券列表，仅包含优惠券基本信息和用户nick
func TmallPromotionCouponQuery(ctx context.Context, clt *core.SDKClient, req *promotion.TmallPromotionCouponQueryAPIRequest, resp *promotion.TmallPromotionCouponQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
