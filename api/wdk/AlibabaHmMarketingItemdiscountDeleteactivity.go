package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaHmMarketingItemdiscountDeleteactivity 删除商品特价活动
// alibaba.hm.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
func AlibabaHmMarketingItemdiscountDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaHmMarketingItemdiscountDeleteactivityAPIRequest, resp *wdk.AlibabaHmMarketingItemdiscountDeleteactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
