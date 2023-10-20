package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountQueryactivity 查找特价活动
// alibaba.hm.marketing.itemdiscount.queryactivity
//
// 查找特价活动
func AlibabaHmMarketingItemdiscountQueryactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountQueryactivityAPIRequest, resp *wdk.AlibabaHmMarketingItemdiscountQueryactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
