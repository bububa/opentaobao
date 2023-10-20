package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// AlibabaMjMoscarnivalReceivecoupon 根据手机号码领券
// alibaba.mj.moscarnival.receivecoupon
//
// 根据手机号码领券
func AlibabaMjMoscarnivalReceivecoupon(clt *core.SDKClient, req *mos.AlibabaMjMoscarnivalReceivecouponAPIRequest, resp *mos.AlibabaMjMoscarnivalReceivecouponAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
