package tbk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/tbk"
)

// TaobaoTbkCartCouponExpireUserQuery 购物车催付优惠券到期查询用户信息
// taobao.tbk.cart.coupon.expire.user.query
//
// 购物车催付根据对应规则查询用户信息。
func TaobaoTbkCartCouponExpireUserQuery(clt *core.SDKClient, req *tbk.TaobaoTbkCartCouponExpireUserQueryAPIRequest, session string) (*tbk.TaobaoTbkCartCouponExpireUserQueryAPIResponse, error) {
	var resp tbk.TaobaoTbkCartCouponExpireUserQueryAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
