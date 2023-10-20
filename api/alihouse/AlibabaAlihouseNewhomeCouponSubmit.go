package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// Alibabaalihousenewhomecouponsubmit 提交专车优惠券活动
// alibaba.alihouse.newhome.coupon.submit
//
// 提交专车优惠券活动
func Alibabaalihousenewhomecouponsubmit(clt *core.SDKClient, req *alihouse.AlibabaalihousenewhomecouponsubmitAPIRequest, session string) (*alihouse.AlibabaalihousenewhomecouponsubmitAPIResponse, error) {
	var resp alihouse.AlibabaalihousenewhomecouponsubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
