package promotion

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallPromotionCouponUserAPIRequest
用户信息查询接口 API请求
tmall.promotion.coupon.user

开发给外部合作商（例如：苏宁），通过会员付款码获得会员nick */
type TmallPromotionCouponUserAPIRequest struct {
	model.Params
	// 例如：suning
	_bizType string
	// 会员付款码
	_payCode string
	// 扩展字段
	_extra string
}

// New
