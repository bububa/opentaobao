package promotion

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/promotion"
)

// AlibabaWdkCouponAbandon 废券
// alibaba.wdk.coupon.abandon
//
// 优惠券废弃
func AlibabaWdkCouponAbandon(clt *core.SDKClient, req *promotion.AlibabaWdkCouponAbandonAPIRequest, session string) (*promotion.AlibabaWdkCouponAbandonAPIResponse, error) {
	var resp promotion.AlibabaWdkCouponAbandonAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
