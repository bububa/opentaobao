package eleenterprisecoupon

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecoupon"
)

// Alibabaeleenterprisecouponsend 发放优惠券
// alibaba.ele.enterprise.coupon.send
//
// 发放优惠券
func Alibabaeleenterprisecouponsend(clt *core.SDKClient, req *eleenterprisecoupon.AlibabaeleenterprisecouponsendAPIRequest, session string) (*eleenterprisecoupon.AlibabaeleenterprisecouponsendAPIResponse, error) {
	var resp eleenterprisecoupon.AlibabaeleenterprisecouponsendAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
