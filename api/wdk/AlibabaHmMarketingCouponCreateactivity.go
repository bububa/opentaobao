package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// Alibabahmmarketingcouponcreateactivity 优惠券活动创建
// alibaba.hm.marketing.coupon.createactivity
//
// 添加优惠券活动
func Alibabahmmarketingcouponcreateactivity(clt *core.SDKClient, req *wdk.AlibabahmmarketingcouponcreateactivityAPIRequest, session string) (*wdk.AlibabahmmarketingcouponcreateactivityAPIResponse, error) {
	var resp wdk.AlibabahmmarketingcouponcreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
