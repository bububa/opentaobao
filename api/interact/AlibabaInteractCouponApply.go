package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

// AlibabaInteractCouponApply 优惠券领取鉴权接口
// alibaba.interact.coupon.apply
//
// 鉴权接口，为coupon.apply接口鉴权
func AlibabaInteractCouponApply(clt *core.SDKClient, req *interact.AlibabaInteractCouponApplyAPIRequest, resp *interact.AlibabaInteractCouponApplyAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
