package eleenterprisecoupon

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecoupon"
)

// Alibabaeleenterprisecartcouponget 获取下单可用的优惠券
// alibaba.ele.enterprise.cartcoupon.get
//
// 获取下单可用的优惠券
func Alibabaeleenterprisecartcouponget(clt *core.SDKClient, req *eleenterprisecoupon.AlibabaeleenterprisecartcoupongetAPIRequest, session string) (*eleenterprisecoupon.AlibabaeleenterprisecartcoupongetAPIResponse, error) {
	var resp eleenterprisecoupon.AlibabaeleenterprisecartcoupongetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
