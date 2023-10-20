package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftCreateactivity 创建买赠活动
// alibaba.hm.marketing.itembuygift.createactivity
//
// 创建买赠活动
func AlibabaHmMarketingItembuygiftCreateactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftCreateactivityAPIRequest, resp *wdk.AlibabaHmMarketingItembuygiftCreateactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
