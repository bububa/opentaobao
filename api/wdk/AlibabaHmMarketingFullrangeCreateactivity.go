package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingFullrangeCreateactivity 创建全场活动
// alibaba.hm.marketing.fullrange.createactivity
//
// 创建全场活动
func AlibabaHmMarketingFullrangeCreateactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingFullrangeCreateactivityAPIRequest, session string) (*wdk.AlibabaHmMarketingFullrangeCreateactivityAPIResponse, error) {
	var resp wdk.AlibabaHmMarketingFullrangeCreateactivityAPIResponse
	err := clt.Post(req, &resp, session)
	if err != nil {
		return nil, err
	}
	return &resp, nil
}
