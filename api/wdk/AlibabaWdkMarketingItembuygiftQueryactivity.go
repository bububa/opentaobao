package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItembuygiftQueryactivity 查询买赠活动
// alibaba.wdk.marketing.itembuygift.queryactivity
//
// 查询买赠活动
func AlibabaWdkMarketingItembuygiftQueryactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftQueryactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItembuygiftQueryactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
