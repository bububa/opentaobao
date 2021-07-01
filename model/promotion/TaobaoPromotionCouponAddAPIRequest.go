package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TaobaoPromotionCouponAddAPIRequest
创建店铺优惠券接口 API请求
taobao.promotion.coupon.add

创建店铺优惠券。有效期内的店铺优惠券总数量不超过50张 */
type TaobaoPromotionCouponAddAPIRequest struct {
	model.Params
	// 优惠券的面额，必须是3，5，10，20，50，100
	_denominations int64
	// 优惠券的截止日期
	_endTime string
	// 订单满多少元才能用这个优惠券，500就是满500元才能使用
	_condition int64
	// 优惠券的生效时间
	_startTime string
}

// New
