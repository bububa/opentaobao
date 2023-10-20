package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingFullrangeCreateactivity 创建全场活动
// alibaba.wdk.marketing.fullrange.createactivity
//
// 创建全场活动
func AlibabaWdkMarketingFullrangeCreateactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeCreateactivityAPIRequest, resp *wdk.AlibabaWdkMarketingFullrangeCreateactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
