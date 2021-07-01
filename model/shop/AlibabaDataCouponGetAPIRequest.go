package shop

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaDataCouponGetAPIRequest
获取优惠券信息 API请求
alibaba.data.coupon.get

获取优惠券信息，仅作客户端鉴权虚拟api使用 */
type AlibabaDataCouponGetAPIRequest struct {
	model.Params
	// 客户端鉴权虚拟api使用
	_unNamed string
}

// New
