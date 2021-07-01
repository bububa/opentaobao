package interact

import (
	"net/url"

	"github.com/bububa/opentaobao/model"
)

/* AlibabaInteractCouponApplyAPIRequest
优惠券领取鉴权接口 API请求
alibaba.interact.coupon.apply

鉴权接口，为coupon.apply接口鉴权 */
type AlibabaInteractCouponApplyAPIRequest struct {
	model.Params
}

// New
