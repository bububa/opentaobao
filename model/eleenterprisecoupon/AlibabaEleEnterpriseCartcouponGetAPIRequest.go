package eleenterprisecoupon

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaEleEnterpriseCartcouponGetAPIRequest
获取下单可用的优惠券 API请求
alibaba.ele.enterprise.cartcoupon.get

获取下单可用的优惠券 */
type AlibabaEleEnterpriseCartcouponGetAPIRequest struct {
	model.Params
	// 手机号
	_phone string
	// 购物车id
	_cartId string
}

// New
