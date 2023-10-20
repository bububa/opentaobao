package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItemdiscountDeleteactivity 删除商品特价活动
// alibaba.wdk.marketing.itemdiscount.deleteactivity
//
// 删除商品特价活动
func AlibabaWdkMarketingItemdiscountDeleteactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountDeleteactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItemdiscountDeleteactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
