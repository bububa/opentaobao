package tmallnr

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* TmallNrtNewcouponSendAPIRequest
券发放接口 API请求
tmall.nrt.newcoupon.send

券发放接口 */
type TmallNrtNewcouponSendAPIRequest struct {
	model.Params
	// 发券dto
	_nrtCouponSendDto *NrtCouponSendDto
}

// New
