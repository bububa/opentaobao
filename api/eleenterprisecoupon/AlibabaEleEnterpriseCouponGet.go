package eleenterprisecoupon

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecoupon"
)

/* AlibabaEleEnterpriseCouponGet
获取用户优惠券
alibaba.ele.enterprise.coupon.get

获取用户优惠券 */
func AlibabaEleEnterpriseCouponGet(clt *core.SDKClient, req *eleenterprisecoupon.AlibabaEleEnterpriseCouponGetAPIRequest, session string) (*eleenterprisecoupon.AlibabaEleEnterpriseCouponGetAPIResponse, error) {
	var resp eleenterprisecoupon.AlibabaEleEnterpriseCouponGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
