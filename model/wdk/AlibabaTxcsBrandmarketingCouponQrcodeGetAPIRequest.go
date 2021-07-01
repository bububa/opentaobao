package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest
品牌营销导购员券页面二维码获取 API请求
alibaba.txcs.brandmarketing.coupon.qrcode.get

构建券页码二维码url */
type AlibabaTxcsBrandmarketingCouponQrcodeGetAPIRequest struct {
	model.Params
	// 请求信息
	_couponQrcodeParamDo *CouponQrcodeParamDo
}

// New
