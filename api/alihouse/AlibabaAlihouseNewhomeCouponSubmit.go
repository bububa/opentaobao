package alihouse

import (
	"context"

	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/alihouse"
)

// AlibabaAlihouseNewhomeCouponSubmit 提交专车优惠券活动
// alibaba.alihouse.newhome.coupon.submit
//
// 提交专车优惠券活动
func AlibabaAlihouseNewhomeCouponSubmit(ctx context.Context, clt *core.SDKClient, req *alihouse.AlibabaAlihouseNewhomeCouponSubmitAPIRequest, resp *alihouse.AlibabaAlihouseNewhomeCouponSubmitAPIResponse, session string) error {
	return clt.Post(ctx, req, resp, session)
}
