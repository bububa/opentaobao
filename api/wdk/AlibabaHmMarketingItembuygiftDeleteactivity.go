package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItembuygiftDeleteactivity 删除买赠活动
// alibaba.hm.marketing.itembuygift.deleteactivity
//
// 删除买赠活动
func AlibabaHmMarketingItembuygiftDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItembuygiftDeleteactivityAPIRequest, resp *wdk.AlibabaHmMarketingItembuygiftDeleteactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
