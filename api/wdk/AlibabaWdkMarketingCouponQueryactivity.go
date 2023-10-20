package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingcouponqueryactivity 查询优惠券活动
// alibaba.wdk.marketing.coupon.queryactivity
//
// 查询优惠券活动
func Alibabawdkmarketingcouponqueryactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingcouponqueryactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingcouponqueryactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingcouponqueryactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
