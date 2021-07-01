package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

/* AlibabaMjMoscarnivalReceivecoupon
根据手机号码领券
alibaba.mj.moscarnival.receivecoupon

根据手机号码领券 */
func AlibabaMjMoscarnivalReceivecoupon(clt *core.SDKClient, req *mos.AlibabaMjMoscarnivalReceivecouponAPIRequest, session string) (*mos.AlibabaMjMoscarnivalReceivecouponAPIResponse, error) {
	var resp mos.AlibabaMjMoscarnivalReceivecouponAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
