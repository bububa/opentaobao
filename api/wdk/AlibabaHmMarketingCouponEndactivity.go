package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingcouponendactivity 结束优惠券活动
// alibaba.hm.marketing.coupon.endactivity
//
// 结束优惠券活动。优惠券变为结束领取状态，已领取的优惠券可以继续使用
func Alibabahmmarketingcouponendactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingcouponendactivityAPIRequest, session string) (*wdk.AlibabahmmarketingcouponendactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingcouponendactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
