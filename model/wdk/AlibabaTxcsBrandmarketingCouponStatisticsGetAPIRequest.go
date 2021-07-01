package wdk

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest
品牌营销导购员券推广统计数据回流 API请求
alibaba.txcs.brandmarketing.coupon.statistics.get

请求券统计数据回流 */
type AlibabaTxcsBrandmarketingCouponStatisticsGetAPIRequest struct {
	model.Params
	// 请求信息
	_couponStatisticsParamDo *CouponStatisticsParamDo
}

// New
