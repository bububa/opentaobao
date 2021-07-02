package eleenterprisecoupon

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/eleenterprisecoupon"
)

// AlibabaEleEnterpriseCartcouponGet 获取下单可用的优惠券
// alibaba.ele.enterprise.cartcoupon.get
//
// 获取下单可用的优惠券
func AlibabaEleEnterpriseCartcouponGet(clt *core.SDKClient, req *eleenterprisecoupon.AlibabaEleEnterpriseCartcouponGetAPIRequest, session string) (*eleenterprisecoupon.AlibabaEleEnterpriseCartcouponGetAPIResponse, error) {
	var resp eleenterprisecoupon.AlibabaEleEnterpriseCartcouponGetAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
