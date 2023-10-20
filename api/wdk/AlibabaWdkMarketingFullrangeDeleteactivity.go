package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingFullrangeDeleteactivity 全场活动删除活动接口
// alibaba.wdk.marketing.fullrange.deleteactivity
//
// 全场活动删除活动
func AlibabaWdkMarketingFullrangeDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingFullrangeDeleteactivityAPIRequest, resp *wdk.AlibabaWdkMarketingFullrangeDeleteactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
