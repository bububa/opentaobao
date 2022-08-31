package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCouponSubmit 提交专车优惠券活动
// alibaba.alihouse.newhome.coupon.submit
//
// 提交专车优惠券活动
func AlibabaAlihouseNewhomeCouponSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCouponSubmitAPIRequest, session string) (*alihouse.AlibabaAlihouseNewhomeCouponSubmitAPIResponse, error) {
	var resp alihouse.AlibabaAlihouseNewhomeCouponSubmitAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
