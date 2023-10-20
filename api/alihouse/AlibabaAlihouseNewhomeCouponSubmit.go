package alihouse

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCouponSubmit 提交专车优惠券活动
// alibaba.alihouse.newhome.coupon.submit
//
// 提交专车优惠券活动
func AlibabaAlihouseNewhomeCouponSubmit(clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCouponSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeCouponSubmitAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
