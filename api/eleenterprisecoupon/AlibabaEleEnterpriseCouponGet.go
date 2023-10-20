package eleenterprisecoupon

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecoupon"
)

// Alibabaeleenterprisecouponget 获取用户优惠券
// alibaba.ele.enterprise.coupon.get
//
// 获取用户优惠券
func Alibabaeleenterprisecouponget(clt *core.SDKClient, req *eleenterprisecoupon.AlibabaeleenterprisecoupongetAPIRequest, session string) (*eleenterprisecoupon.AlibabaeleenterprisecoupongetAPIResponse, error) {
	var resp eleenterprisecoupon.AlibabaeleenterprisecoupongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
