package tbk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoTbkCartCouponExpireUserQueryAPIRequest
购物车催付优惠券到期查询用户信息 API请求
taobao.tbk.cart.coupon.expire.user.query

购物车催付根据对应规则查询用户信息。 */
type TaobaoTbkCartCouponExpireUserQueryAPIRequest struct {
	model.Params
	// 规则ID，由接口提供方分配
	_ruleId string
	// 每页大小
	_pageSize int64
	// 页码，从0开始
	_pageNum int64
}

// New
