package nrt

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtCouponSendAPIRequest
券发放接口 API请求
tmall.nrt.coupon.send

新零售场景，商家自有渠道发放券 */
type TmallNrtCouponSendAPIRequest struct {
	model.Params
	// 发券dto
	_nrtCouponSendDto *NrtCouponSendDto
}

// New
