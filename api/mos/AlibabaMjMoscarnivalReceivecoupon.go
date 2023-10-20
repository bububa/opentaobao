package mos

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/mos"
)

// Alibabamjmoscarnivalreceivecoupon 根据手机号码领券
// alibaba.mj.moscarnival.receivecoupon
//
// 根据手机号码领券
func Alibabamjmoscarnivalreceivecoupon(clt *core.SDKClient, req *mos.AlibabamjmoscarnivalreceivecouponAPIRequest, session string) (*mos.AlibabamjmoscarnivalreceivecouponAPIResponse, error) {
	var resp mos.AlibabamjmoscarnivalreceivecouponAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
