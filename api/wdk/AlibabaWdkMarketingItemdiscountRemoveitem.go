package wdk

import (
	"github.com/bububa/opentaobao/core"
	"github.com/bububa/opentaobao/model/wdk"
)

// AlibabaWdkMarketingItemdiscountRemoveitem 移除报名的商品
// alibaba.wdk.marketing.itemdiscount.removeitem
//
// 将报名特价活动的商品从特价活动中移除
func AlibabaWdkMarketingItemdiscountRemoveitem(clt *core.SDKClient, req *wdk.AlibabaWdkMarketingItemdiscountRemoveitemAPIRequest, resp *wdk.AlibabaWdkMarketingItemdiscountRemoveitemAPIResponse, session string) error {
	return clt.Post(req, resp, session)
}
