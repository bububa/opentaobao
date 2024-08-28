package tbk

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkCartCouponExpireUserQuery 购物车催付优惠券到期查询用户信息
// taobao.tbk.cart.coupon.expire.user.query
//
// 购物车催付根据对应规则查询用户信息。
func TaobaoTbkCartCouponExpireUserQuery(ctx context.Context, clt *core.SDKClient, req *tbk.TaobaoTbkCartCouponExpireUserQueryAPIRequest, resp *tbk.TaobaoTbkCartCouponExpireUserQueryAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
