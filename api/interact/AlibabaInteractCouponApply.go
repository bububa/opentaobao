package interact

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/interact"
)

/* AlibabaInteractCouponApply
优惠券领取鉴权接口
alibaba.interact.coupon.apply

鉴权接口，为coupon.apply接口鉴权 */
func AlibabaInteractCouponApply(clt *core.SDKClient, req *interact.AlibabaInteractCouponApplyAPIRequest, session string) (*interact.AlibabaInteractCouponApplyAPIResponse, error) {
	var resp interact.AlibabaInteractCouponApplyAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
