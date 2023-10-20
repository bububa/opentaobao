package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItemdiscountCreateactivity 创建商品特价活动
// alibaba.wdk.marketing.itemdiscount.createactivity
//
// 创建商品特价活动
func AlibabaWdkMarketingItemdiscountCreateactivity(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountCreateactivityAPIRequest, resp *wdk.AlibabaWdkMarketingItemdiscountCreateactivityAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
