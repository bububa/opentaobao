package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabawdkmarketingcouponcreateactivity 优惠券活动创建
// alibaba.wdk.marketing.coupon.createactivity
//
// 添加优惠券活动
func Alibabawdkmarketingcouponcreateactivity(clt *core.SDKClient, req *wdk.AlibabawdkmarketingcouponcreateactivityAPIRequest, session string) (*wdk.AlibabawdkmarketingcouponcreateactivityAPIResponse, error) {
	var resp wdk.AlibabawdkmarketingcouponcreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
