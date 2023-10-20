package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingCouponCreateactivity 优惠券活动创建
// alibaba.wdk.marketing.coupon.createactivity
//
// 添加优惠券活动
func AlibabaWdkMarketingCouponCreateactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingCouponCreateactivityAPIRequest, session string) (*wdk.AlibabaWdkMarketingCouponCreateactivityAPIResponse, error) {
	var resp wdk.AlibabaWdkMarketingCouponCreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
