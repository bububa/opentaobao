package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItembuygiftDeleteactivity 删除买赠活动
// alibaba.wdk.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
func AlibabaWdkMarketingItembuygiftDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItembuygiftDeleteactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItembuygiftDeleteactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
